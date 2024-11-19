package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestTrackHandlers_SearchTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful search", func(t *testing.T) {
		tracks := []*dto.TrackDTO{
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album1",
			},
			{
				Name: "track", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist2", Album: "album2",
			},
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist3", Album: "album3",
			},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "test").Return([]*dto.TrackDTO{tracks[0], tracks[2]}, nil)

		request, err := http.NewRequest(http.MethodGet, "/tracks/search/?name=test", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		trackHandlers.SearchTrack(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundTracks []*dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&foundTracks)
		assert.NoError(t, err)

		expectedTracks := []*dto.TrackDTO{tracks[0], tracks[2]}
		assert.Equal(t, expectedTracks, foundTracks)
	})

	t.Run("Missing query parameter", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/tracks/search/", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		trackHandlers.SearchTrack(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Can't find tracks", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/tracks/search/?name=song", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "song").Return([]*dto.TrackDTO{}, nil)

		trackHandlers.SearchTrack(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestTrackHandlers_ViewTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful view", func(t *testing.T) {
		track := dto.TrackDTO{
			Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
			Artist: "artist1", Album: "album1",
		}

		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(&track, nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/{id}", trackHandlers.ViewTrack).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/tracks/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundTrack dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&foundTrack)
		assert.NoError(t, err)

		assert.Equal(t, track, foundTrack)
	})

	t.Run("Wrong slug", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/tracks/{id}", trackHandlers.ViewTrack).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/tracks/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Can't find track", func(t *testing.T) {
		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(nil, errors.New("Can't find track"))

		router := mux.NewRouter()
		router.HandleFunc("/tracks/{id}", trackHandlers.ViewTrack).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/tracks/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func TestTrackHandlers_GetAllTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful got all tracks", func(t *testing.T) {
		tracks := []*dto.TrackDTO{
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album1",
			},
			{
				Name: "track", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist2", Album: "album2",
			},
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist3", Album: "album3",
			},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return(tracks, nil)

		request, err := http.NewRequest(http.MethodGet, "/tracks", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		trackHandlers.GetAll(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundTracks []*dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&foundTracks)
		assert.NoError(t, err)

		assert.Equal(t, tracks, foundTracks)
	})

	t.Run("Can't find tracks", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/tracks", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return([]*dto.TrackDTO{}, nil)

		trackHandlers.GetAll(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestTrackHandlers_GetAllByArtistIDTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful got all tracks by artist ID", func(t *testing.T) {
		tracks := []*dto.TrackDTO{
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album1",
			},
			{
				Name: "track", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album2",
			},
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album3",
			},
		}
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return(tracks, nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/byArtistId/{artistId}", trackHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundTracks []*dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&foundTracks)
		assert.NoError(t, err)

		assert.Equal(t, tracks, foundTracks)
	})

	t.Run("Can't find tracks by artist ID", func(t *testing.T) {
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return([]*dto.TrackDTO{}, nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/byArtistId/{artistId}", trackHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/tracks/byArtistId/{artistId}", trackHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byArtistId/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestTrackHandlers_GetAllByAlbumIDTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful got all tracks by album ID", func(t *testing.T) {
		tracks := []*dto.TrackDTO{
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist1", Album: "album1",
			},
			{
				Name: "track", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist2", Album: "album1",
			},
			{
				Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
				Artist: "artist3", Album: "album1",
			},
		}
		usecaseMock.EXPECT().GetAllByAlbumID(gomock.Any(), uint64(1)).Return(tracks, nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/byAlbumId/{albumId}", trackHandlers.GetAllByAlbumID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byAlbumId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundTracks []*dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&foundTracks)
		assert.NoError(t, err)

		assert.Equal(t, tracks, foundTracks)
	})

	t.Run("Can't find tracks by album ID", func(t *testing.T) {
		usecaseMock.EXPECT().GetAllByAlbumID(gomock.Any(), uint64(1)).Return([]*dto.TrackDTO{}, nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/byAlbumId/{albumId}", trackHandlers.GetAllByAlbumID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byAlbumId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Invalid album ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/tracks/byAlbumId/{albumId}", trackHandlers.GetAllByAlbumID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/tracks/byAlbumId/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestTrackHandlers_AddFavoriteTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful add track to favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		usecaseMock.EXPECT().AddFavoriteTrack(gomock.Any(), userID, trackID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.AddFavoriteTrack).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid track ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.AddFavoriteTrack).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, "/tracks/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		trackID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.AddFavoriteTrack).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error in usecase when adding track to favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().AddFavoriteTrack(gomock.Any(), userID, trackID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.AddFavoriteTrack).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't add track to favorite")
	})
}

func TestTrackHandlers_DeleteFavoriteTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Successful delete track from favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		usecaseMock.EXPECT().DeleteFavoriteTrack(gomock.Any(), userID, trackID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.DeleteFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid track ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.DeleteFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, "/tracks/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		trackID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.DeleteFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error in usecase when deleting track from favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().DeleteFavoriteTrack(gomock.Any(), userID, trackID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/tracks/favorite/{trackID}", trackHandlers.DeleteFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't delete track from favorite")
	})
}

func TestTrackHandlers_IsFavoriteTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Track is in favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		usecaseMock.EXPECT().IsFavoriteTrack(gomock.Any(), userID, trackID).Return(true, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite/{trackID}", trackHandlers.IsFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result map[string]bool
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.True(t, result["exists"])
	})

	t.Run("Track is not in favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		usecaseMock.EXPECT().IsFavoriteTrack(gomock.Any(), userID, trackID).Return(false, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite/{trackID}", trackHandlers.IsFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result map[string]bool
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.False(t, result["exists"])
	})

	t.Run("Invalid track ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite/{trackID}", trackHandlers.IsFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, "/api/v1/tracks/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		trackID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite/{trackID}", trackHandlers.IsFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error when checking if track is in favorites", func(t *testing.T) {
		userID := uuid.New()
		trackID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().IsFavoriteTrack(gomock.Any(), userID, trackID).Return(false, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite/{trackID}", trackHandlers.IsFavoriteTrack).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tracks/favorite/%d", trackID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't check is track in favorite")
	})
}

func TestTrackHandlers_GetFavoriteTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	trackHandlers := NewTrackHandlers(usecaseMock, logger)

	t.Run("Success", func(t *testing.T) {
		userID := uuid.New()
		tracks := []*dto.TrackDTO{
			{ID: 1, Name: "Track 1", Artist: "Artist 1"},
			{ID: 2, Name: "Track 2", Artist: "Artist 2"},
		}

		usecaseMock.EXPECT().GetFavoriteTracks(gomock.Any(), userID).Return(tracks, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite", trackHandlers.GetFavoriteTracks).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/tracks/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result []*dto.TrackDTO
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Track 1", result[0].Name)
		assert.Equal(t, "Track 2", result[1].Name)
	})

	t.Run("User ID not found", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite", trackHandlers.GetFavoriteTracks).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/tracks/favorite", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error while getting favorite tracks", func(t *testing.T) {
		userID := uuid.New()
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetFavoriteTracks(gomock.Any(), userID).Return(nil, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite", trackHandlers.GetFavoriteTracks).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/tracks/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Failed to get favorite tracks")
	})

	t.Run("No favorite tracks found", func(t *testing.T) {
		userID := uuid.New()
		usecaseMock.EXPECT().GetFavoriteTracks(gomock.Any(), userID).Return(nil, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/tracks/favorite", trackHandlers.GetFavoriteTracks).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/tracks/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Contains(t, response.Body.String(), "No favorite tracks were found")
	})
}
