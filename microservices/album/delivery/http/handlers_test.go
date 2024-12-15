package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAlbumHandlers_SearchAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful search", func(t *testing.T) {
		releaseDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
		albums := []*dto.AlbumDTO{
			{
				Name: "test", ReleaseDate: releaseDate, Image: "1", ArtistName: "1",
			},
			{
				Name: "test", ReleaseDate: releaseDate, Image: "2", ArtistName: "2",
			},
			{
				Name: "album", ReleaseDate: releaseDate, Image: "3", ArtistName: "3",
			},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "test").Return([]*dto.AlbumDTO{albums[0], albums[1]}, nil)

		request, err := http.NewRequest(http.MethodGet, "/albums/search/?query=test", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		albumHandlers.SearchAlbum(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundAlbums []*dto.AlbumDTO
		err = json.NewDecoder(res.Body).Decode(&foundAlbums)
		assert.NoError(t, err)

		expectedAlbums := []*dto.AlbumDTO{albums[0], albums[1]}
		assert.Equal(t, expectedAlbums, foundAlbums)
	})

	t.Run("Missing query parameter", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/albums/search/", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		albumHandlers.SearchAlbum(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Can't find albums", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/albums/search/?query=album", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().Search(ctx, "album").Return([]*dto.AlbumDTO{}, nil)

		albumHandlers.SearchAlbum(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestAlbumHandlers_ViewAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful view", func(t *testing.T) {
		releaseDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
		album := dto.AlbumDTO{
			Name: "test", ReleaseDate: releaseDate, Image: "1", ArtistName: "1",
		}

		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(&album, nil)

		router := mux.NewRouter()
		router.HandleFunc("/albums/{id}", albumHandlers.ViewAlbum).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/albums/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundAlbum dto.AlbumDTO
		err = json.NewDecoder(res.Body).Decode(&foundAlbum)
		assert.NoError(t, err)

		assert.Equal(t, album, foundAlbum)
	})

	t.Run("Wrong slug", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/albums/{id}", albumHandlers.ViewAlbum).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/albums/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Can't find album", func(t *testing.T) {
		usecaseMock.EXPECT().View(gomock.Any(), uint64(1)).Return(nil, errors.New("Can't find album"))

		router := mux.NewRouter()
		router.HandleFunc("/albums/{id}", albumHandlers.ViewAlbum).Methods("GET")
		request, err := http.NewRequest(http.MethodGet, "/albums/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func TestAlbumHandlers_GetAllAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful got all albums", func(t *testing.T) {
		releaseDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
		albums := []*dto.AlbumDTO{
			{
				Name: "test", ReleaseDate: releaseDate, Image: "1", ArtistName: "1",
			},
			{
				Name: "test", ReleaseDate: releaseDate, Image: "2", ArtistName: "2",
			},
			{
				Name: "album", ReleaseDate: releaseDate, Image: "3", ArtistName: "3",
			},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return(albums, nil)

		request, err := http.NewRequest(http.MethodGet, "/albums", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		albumHandlers.GetAll(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundAlbums []*dto.AlbumDTO
		err = json.NewDecoder(res.Body).Decode(&foundAlbums)
		assert.NoError(t, err)

		assert.Equal(t, albums, foundAlbums)
	})

	t.Run("Can't find albums", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/albums", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return([]*dto.AlbumDTO{}, nil)

		albumHandlers.GetAll(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestAlbumHandlers_GetAllByArtistIDAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful got all albums by artist ID", func(t *testing.T) {
		releaseDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
		albums := []*dto.AlbumDTO{
			{
				Name: "test", ReleaseDate: releaseDate, Image: "1", ArtistName: "artist1",
			},
			{
				Name: "album", ReleaseDate: releaseDate, Image: "2", ArtistName: "artist1",
			},
			{
				Name: "test", ReleaseDate: releaseDate, Image: "3", ArtistName: "artist1",
			},
		}
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return(albums, nil)

		router := mux.NewRouter()
		router.HandleFunc("/albums/byArtistId/{artistId}", albumHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/albums/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundAlbums []*dto.AlbumDTO
		err = json.NewDecoder(res.Body).Decode(&foundAlbums)
		assert.NoError(t, err)

		assert.Equal(t, albums, foundAlbums)
	})

	t.Run("Can't find albums by artist ID", func(t *testing.T) {
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return([]*dto.AlbumDTO{}, nil)

		router := mux.NewRouter()
		router.HandleFunc("/albums/byArtistId/{artistId}", albumHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/albums/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/albums/byArtistId/{artistId}", albumHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/albums/byArtistId/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestAlbumHandlers_AddFavoriteAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful add album to favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		usecaseMock.EXPECT().AddFavoriteAlbum(gomock.Any(), userID, albumID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.AddFavoriteAlbum).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid album ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.AddFavoriteAlbum).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, "/albums/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		albumID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.AddFavoriteAlbum).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error in usecase when adding album to favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().AddFavoriteAlbum(gomock.Any(), userID, albumID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.AddFavoriteAlbum).Methods("POST")

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't add album to favorite")
	})
}

func TestAlbumHandlers_DeleteFavoriteAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Successful delete album from favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		usecaseMock.EXPECT().DeleteFavoriteAlbum(gomock.Any(), userID, albumID).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.DeleteFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Invalid album ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.DeleteFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, "/albums/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Error in usecase when deleting album from favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().DeleteFavoriteAlbum(gomock.Any(), userID, albumID).Return(mockError)

		router := mux.NewRouter()
		router.HandleFunc("/albums/favorite/{albumID}", albumHandlers.DeleteFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't delete album from favorite")
	})
}

func TestAlbumHandlers_IsFavoriteAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Album is in favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		usecaseMock.EXPECT().IsFavoriteAlbum(gomock.Any(), userID, albumID).Return(true, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite/{albumID}", albumHandlers.IsFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/albums/favorite/%d", albumID), nil)
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

	t.Run("Album is not in favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		usecaseMock.EXPECT().IsFavoriteAlbum(gomock.Any(), userID, albumID).Return(false, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite/{albumID}", albumHandlers.IsFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/albums/favorite/%d", albumID), nil)
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

	t.Run("Invalid album ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite/{albumID}", albumHandlers.IsFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, "/api/v1/albums/favorite/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("User ID not found in context", func(t *testing.T) {
		albumID := uint64(1)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite/{albumID}", albumHandlers.IsFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error when checking if album is in favorites", func(t *testing.T) {
		userID := uuid.New()
		albumID := uint64(1)
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().IsFavoriteAlbum(gomock.Any(), userID, albumID).Return(false, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite/{albumID}", albumHandlers.IsFavoriteAlbum).Methods("DELETE")

		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/albums/favorite/%d", albumID), nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Can't check is album in favorite")
	})
}

func TestAlbumHandlers_GetFavoriteAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	albumHandlers := NewAlbumHandlers(usecaseMock, logger)

	t.Run("Success", func(t *testing.T) {
		userID := uuid.New()
		albums := []*dto.AlbumDTO{
			{ID: 1, Name: "Album 1", ArtistName: "Artist 1"},
			{ID: 2, Name: "Album 2", ArtistName: "Artist 2"},
		}

		usecaseMock.EXPECT().GetFavoriteAlbums(gomock.Any(), userID).Return(albums, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite", albumHandlers.GetFavoriteAlbums).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/albums/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var result []*dto.AlbumDTO
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Album 1", result[0].Name)
		assert.Equal(t, "Album 2", result[1].Name)
	})

	t.Run("User ID not found", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite", albumHandlers.GetFavoriteAlbums).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/albums/favorite", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, response.Body.String(), "User id not found")
	})

	t.Run("Error while getting favorite albums", func(t *testing.T) {
		userID := uuid.New()
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetFavoriteAlbums(gomock.Any(), userID).Return(nil, mockError)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite", albumHandlers.GetFavoriteAlbums).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/albums/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, response.Body.String(), "Failed to get favorite albums")
	})

	t.Run("No favorite albums found", func(t *testing.T) {
		userID := uuid.New()
		usecaseMock.EXPECT().GetFavoriteAlbums(gomock.Any(), userID).Return(nil, nil)

		router := mux.NewRouter()
		router.HandleFunc("/api/v1/albums/favorite", albumHandlers.GetFavoriteAlbums).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/api/v1/albums/favorite", nil)
		assert.NoError(t, err)
		request = request.WithContext(context.WithValue(request.Context(), utils.UserIDKey{}, userID))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Contains(t, response.Body.String(), "No favorite albums were found")
	})
}
