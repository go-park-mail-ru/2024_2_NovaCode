package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
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
