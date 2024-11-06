package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/dto"
	mocks "github.com/go-park-mail-ru/2024_2_NovaCode/internal/track/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/golang/mock/gomock"
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
