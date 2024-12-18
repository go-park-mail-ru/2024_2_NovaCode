package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestArtistHandlers_SearchArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful search", func(t *testing.T) {
		artists := []*dto.ArtistDTO{
			{Name: "test", Bio: "1", Country: "1", Image: "1"},
			{Name: "artist", Bio: "2", Country: "2", Image: "2"},
			{Name: "artist", Bio: "3", Country: "3", Image: "3"},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "artist").Return([]*dto.ArtistDTO{artists[1], artists[2]}, nil)

		request, err := http.NewRequest(http.MethodGet, "/artists/search/?query=artist", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		artistHandlers.SearchArtist(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundArtists []*dto.ArtistDTO
		err = json.NewDecoder(res.Body).Decode(&foundArtists)
		assert.NoError(t, err)

		expectedArtists := []*dto.ArtistDTO{artists[1], artists[2]}
		assert.Equal(t, expectedArtists, foundArtists)
	})

	t.Run("Missing query parameter", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/artists/search/", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		artistHandlers.SearchArtist(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Can't find artists", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/artists/search/?query=artist", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "artist").Return([]*dto.ArtistDTO{}, nil)

		artistHandlers.SearchArtist(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestArtistHandlers_ViewArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful view", func(t *testing.T) {
		artist := dto.ArtistDTO{Name: "test", Bio: "1", Country: "1", Image: "1"}

		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(&artist, nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/{id}", artistHandlers.ViewArtist).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/artists/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundArtist dto.ArtistDTO
		err = json.NewDecoder(res.Body).Decode(&foundArtist)
		assert.NoError(t, err)

		assert.Equal(t, artist, foundArtist)
	})

	t.Run("Wrong slug", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/artists/{id}", artistHandlers.ViewArtist).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/artists/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Can't find artist", func(t *testing.T) {
		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(nil, errors.New("Can't find artist"))

		router := mux.NewRouter()
		router.HandleFunc("/artists/{id}", artistHandlers.ViewArtist).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/artists/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func TestArtistHandlers_GetAllArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful got all artists", func(t *testing.T) {
		artists := []*dto.ArtistDTO{
			{Name: "test", Bio: "1", Country: "1", Image: "1"},
			{Name: "artist", Bio: "2", Country: "2", Image: "2"},
			{Name: "artist", Bio: "3", Country: "3", Image: "3"},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return(artists, nil)

		request, err := http.NewRequest(http.MethodGet, "/artists", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		artistHandlers.GetAll(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundArtists []*dto.ArtistDTO
		err = json.NewDecoder(res.Body).Decode(&foundArtists)
		assert.NoError(t, err)

		assert.Equal(t, artists, foundArtists)
	})

	t.Run("Can't find artists", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/artists", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return([]*dto.ArtistDTO{}, nil)

		artistHandlers.GetAll(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestArtistHandlers_AddFavoriteArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful add artist to favorites", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		usecaseMock.EXPECT().AddFavoriteArtist(gomock.Any(), userID, artistID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.AddFavoriteArtist).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.AddFavoriteArtist).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, "/artists/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		artistID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.AddFavoriteArtist).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error in usecase when adding artist to favorites", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().AddFavoriteArtist(gomock.Any(), userID, artistID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.AddFavoriteArtist).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't add artist to favorite")
	})
}

func TestArtistHandlers_DeleteFavoriteArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful delete artist from favorites", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		usecaseMock.EXPECT().DeleteFavoriteArtist(gomock.Any(), userID, artistID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.DeleteFavoriteArtist).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.DeleteFavoriteArtist).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, "/artists/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Error in usecase when deleting artist from favorites", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().DeleteFavoriteArtist(gomock.Any(), userID, artistID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.DeleteFavoriteArtist).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't delete artist from favorite")
	})
}

func TestArtistHandlers_IsFavoriteArtist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful check if artist is favorite", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		usecaseMock.EXPECT().IsFavoriteArtist(gomock.Any(), userID, artistID).Return(true, nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.IsFavoriteArtist).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, response.Body.String(), `"exists":true`)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.IsFavoriteArtist).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/artists/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Error in usecase", func(t *testing.T) {
		userID := uuid.New()
		artistID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().IsFavoriteArtist(gomock.Any(), userID, artistID).Return(false, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite/{artistID}", artistHandlers.IsFavoriteArtist).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/artists/favorite/%d", artistID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

func TestArtistHandlers_GetFavoriteArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Successful retrieval of favorite artists", func(t *testing.T) {
		userID := uuid.New()
		artists := []*dto.ArtistDTO{
			{ID: 1, Name: "Artist1"},
			{ID: 2, Name: "Artist2"},
		}

		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(artists, nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorites", artistHandlers.GetFavoriteArtists).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/artists/favorites", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
		var result []dto.ArtistDTO
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, len(artists), len(result))
	})

	t.Run("No favorite artists found", func(t *testing.T) {
		userID := uuid.New()
		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(nil, nil)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite", artistHandlers.GetFavoriteArtists).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/artists/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Error in usecase when retrieving favorite artists", func(t *testing.T) {
		userID := uuid.New()
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(nil, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/artists/favorite", artistHandlers.GetFavoriteArtists).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/artists/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

func TestArtistHandlers_GetFavoriteArtistsCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	artistHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Success", func(t *testing.T) {
		userID := uuid.New()
		count := uint64(2)

		usecaseMock.EXPECT().GetFavoriteArtistsCount(gomock.Any(), userID).Return(count, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/favorite/count/{userID:[0-9a-fA-F-]+}", artistHandlers.GetFavoriteArtistsCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/artists/favorite/count/%s", userID.String()), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result map[string]uint64
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, count, result["favoriteArtistsCount"])
	})

	t.Run("Wrong id value", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/favorite/count/{userID:[0-9a-fA-F-]+}", artistHandlers.GetFavoriteArtistsCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/artists/favorite/count/123", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Wrong id value")
	})

	t.Run("Error while getting favorite artists count", func(t *testing.T) {
		userID := uuid.New()
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetFavoriteArtistsCount(gomock.Any(), userID).Return(uint64(0), mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/favorite/count/{userID:[0-9a-fA-F-]+}", artistHandlers.GetFavoriteArtistsCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/artists/favorite/count/%s", userID.String()), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Failed to get favorite artists count")
	})

	t.Run("No favorite artists found", func(t *testing.T) {
		userID := uuid.New()
		usecaseMock.EXPECT().GetFavoriteArtistsCount(gomock.Any(), userID).Return(uint64(0), nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/favorite/count/{userID:[0-9a-fA-F-]+}", artistHandlers.GetFavoriteArtistsCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/artists/favorite/count/%s", userID.String()), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Contains(t, response.Body.String(), "No favorite artists were found")
	})
}

func TestAlbumHandlers_GetArtistLikesCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewArtistHandlers(usecaseMock, logger)

	t.Run("Success", func(t *testing.T) {
		artistID := uint64(123)
		likesCount := uint64(10)

		usecaseMock.EXPECT().GetArtistLikesCount(gomock.Any(), artistID).Return(likesCount, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/likes/{artistID:[0-9]+}", albumHandlers.GetArtistLikesCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/artists/likes/%d", artistID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result map[string]uint64
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, likesCount, result["artistLikesCount"])
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/likes/{artistID}", albumHandlers.GetArtistLikesCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/artists/likes/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Invalid artist ID")
	})

	t.Run("Error while getting artist likes count", func(t *testing.T) {
		artistID := uint64(123)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetArtistLikesCount(gomock.Any(), artistID).Return(uint64(0), mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/artists/likes/{artistID:[0-9]+}", albumHandlers.GetArtistLikesCount).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/artists/likes/%d", artistID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't get artist likes count")
	})
}
