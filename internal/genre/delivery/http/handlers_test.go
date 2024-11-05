package http

import (
	"context"
	"encoding/json"
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGenreHandlers_GetAllGenres(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	genreHandlers := NewGenreHandlers(usecaseMock, logger)

	t.Run("Successful got all genres", func(t *testing.T) {
		genres := []*dto.GenreDTO{
			{
				ID:      1,
				Name:    "Rock",
				RusName: "Рок",
			},
			{
				ID:      2,
				Name:    "Pop",
				RusName: "Поп",
			},
			{
				ID:      3,
				Name:    "Jazz",
				RusName: "Джаз",
			},
		}

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return(genres, nil)

		request, err := http.NewRequest(http.MethodGet, "/genres", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()

		genreHandlers.GetAll(response, request)
		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundGenres []*dto.GenreDTO
		err = json.NewDecoder(res.Body).Decode(&foundGenres)
		assert.NoError(t, err)

		assert.Equal(t, genres, foundGenres)
	})

	t.Run("Can't find genres", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/genres", nil)
		assert.NoError(t, err)
		response := httptest.NewRecorder()

		ctx := context.Background()
		usecaseMock.EXPECT().GetAll(ctx).Return([]*dto.GenreDTO{}, nil)

		genreHandlers.GetAll(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestGenreHandlers_GetAllByArtistIDGenres(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	genreHandlers := NewGenreHandlers(usecaseMock, logger)

	t.Run("Successful got all genres by artist ID", func(t *testing.T) {
		genres := []*dto.GenreDTO{
			{
				ID:      1,
				Name:    "Rock",
				RusName: "Рок",
			},
			{
				ID:      2,
				Name:    "Pop",
				RusName: "Поп",
			},
			{
				ID:      3,
				Name:    "Jazz",
				RusName: "Джаз",
			},
		}
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return(genres, nil)

		router := mux.NewRouter()
		router.HandleFunc("/genres/byArtistId/{artistId}", genreHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundGenres []*dto.GenreDTO
		err = json.NewDecoder(res.Body).Decode(&foundGenres)
		assert.NoError(t, err)

		assert.Equal(t, genres, foundGenres)
	})

	t.Run("Can't find genres by artist ID", func(t *testing.T) {
		usecaseMock.EXPECT().GetAllByArtistID(gomock.Any(), uint64(1)).Return([]*dto.GenreDTO{}, nil)

		router := mux.NewRouter()
		router.HandleFunc("/genres/byArtistId/{artistId}", genreHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byArtistId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Invalid artist ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/genres/byArtistId/{artistId}", genreHandlers.GetAllByArtistID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byArtistId/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestGenreHandlers_GetAllByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{}
	logger := logger.New(&cfg.Service.Logger)
	usecaseMock := mocks.NewMockUsecase(ctrl)
	genreHandlers := NewGenreHandlers(usecaseMock, logger)

	t.Run("Successful got all genres by track ID", func(t *testing.T) {
		genres := []*dto.GenreDTO{
			{
				ID:      1,
				Name:    "Rock",
				RusName: "Рок",
			},
			{
				ID:      2,
				Name:    "Pop",
				RusName: "Поп",
			},
			{
				ID:      3,
				Name:    "Jazz",
				RusName: "Джаз",
			},
		}
		usecaseMock.EXPECT().GetAllByTrackID(gomock.Any(), uint64(1)).Return(genres, nil)

		router := mux.NewRouter()
		router.HandleFunc("/genres/byTrackId/{trackId}", genreHandlers.GetAllByTrackID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byTrackId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		defer res.Body.Close()
		var foundGenres []*dto.GenreDTO
		err = json.NewDecoder(res.Body).Decode(&foundGenres)
		assert.NoError(t, err)

		assert.Equal(t, genres, foundGenres)
	})

	t.Run("Can't find genres by track ID", func(t *testing.T) {
		usecaseMock.EXPECT().GetAllByTrackID(gomock.Any(), uint64(1)).Return([]*dto.GenreDTO{}, nil)

		router := mux.NewRouter()
		router.HandleFunc("/genres/byTrackId/{trackId}", genreHandlers.GetAllByTrackID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byTrackId/1", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Invalid track ID", func(t *testing.T) {
		router := mux.NewRouter()
		router.HandleFunc("/genres/byTrackId/{trackId}", genreHandlers.GetAllByTrackID).Methods("GET")

		request, err := http.NewRequest(http.MethodGet, "/genres/byTrackId/abc", nil)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := response.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
