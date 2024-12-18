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

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/artists/favorite/byUser/{userID}", artistHandlers.GetFavoriteArtists).Methods("GET")

	t.Run("Successful retrieval of favorite artists", func(t *testing.T) {
		requestID := "test-request-id"
		userID := uuid.New()
		artists := []*dto.ArtistDTO{
			{ID: 1, Name: "Artist1"},
			{ID: 2, Name: "Artist2"},
		}

		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(artists, nil)

		request := createArtistRequestWithVars(requestID, userID)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)

		var result []*dto.ArtistDTO
		err := json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Artist1", result[0].Name)
		assert.Equal(t, "Artist2", result[1].Name)
	})

	t.Run("No favorite artists found", func(t *testing.T) {
		requestID := "test-request-id"
		userID := uuid.New()
		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(nil, nil)

		request := createArtistRequestWithVars(requestID, userID)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Contains(t, response.Body.String(), "No favorite artists were found")
	})

	t.Run("Invalid user ID", func(t *testing.T) {
		requestID := "test-request-id"
		invalidUserID := "invalid-uuid"

		request := httptest.NewRequest(http.MethodGet, "/api/v1/artists/favorite/byUser/"+invalidUserID, nil)
		request = request.WithContext(context.WithValue(request.Context(), utils.RequestIDKey{}, requestID))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), "Invalid user ID")
	})

	t.Run("Error in usecase when retrieving favorite artists", func(t *testing.T) {
		requestID := "test-request-id"
		userID := uuid.New()
		mockError := fmt.Errorf("usecase error")
		usecaseMock.EXPECT().GetFavoriteArtists(gomock.Any(), userID).Return(nil, mockError)

		request := createArtistRequestWithVars(requestID, userID)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Contains(t, response.Body.String(), "Failed to get favorite artists")
	})
}

func createArtistRequestWithVars(requestID string, userID uuid.UUID) *http.Request {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/artists/favorite/byUser/"+userID.String(), nil)
	request = request.WithContext(context.WithValue(request.Context(), utils.RequestIDKey{}, requestID))
	return request
}
