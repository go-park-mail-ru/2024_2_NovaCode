package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/mock"
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

		request, err := http.NewRequest(http.MethodGet, "/artists/search/?name=artist", nil)
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
		request, err := http.NewRequest(http.MethodGet, "/artists/search/?name=artist", nil)
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
