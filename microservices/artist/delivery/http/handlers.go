package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type artistHandlers struct {
	usecase artist.Usecase
	logger  logger.Logger
}

func NewArtistHandlers(usecase artist.Usecase, logger logger.Logger) artist.Handlers {
	return &artistHandlers{usecase, logger}
}

// SearchArtist godoc
// @Summary Search artists by query
// @Description Searches for artists based on the provided "query" parameter.
// @Param query query string true "Name of the artist to search for"
// @Success 200 {array}  dto.ArtistDTO "List of found artists"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No artists found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode artists"
// @Router /api/v1/artists/search [get]
func (handlers *artistHandlers) SearchArtist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	query := request.URL.Query().Get("query")
	if query == "" {
		handlers.logger.Error("Missing query parameter 'query'", requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundArtists, err := handlers.usecase.Search(request.Context(), query)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't find artists")
		return
	} else if len(foundArtists) == 0 {
		handlers.logger.Error(fmt.Sprintf("No artists with %s were found", query), requestID)
		utils.JSONError(response, http.StatusNotFound, "No artists")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// ViewArtist godoc
// @Summary Get artist by ID
// @Description Retrieves an artist using the provided artist ID in the URL path.
// @Param id path uint64 true "Artist ID"
// @Success 200 {object} dto.ArtistDTO "Artist found"
// @Failure 400 {object} utils.ErrorResponse "Invalid artist ID"
// @Failure 404 {object} utils.ErrorResponse "Artist not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to encode the artist data"
// @Router /api/v1/artists/{id} [get]
func (handlers *artistHandlers) ViewArtist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundArtist, err := handlers.usecase.View(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Arist wasn't found: %v", err), requestID)
		utils.JSONError(response, http.StatusNotFound, "Can't find artist")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtist); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artist: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetAll godoc
// @Summary Get all artists
// @Description Retrieves a list of all artists from the database.
// @Success 200 {array} dto.ArtistDTO "List of all artists"
// @Failure 404 {object} utils.ErrorResponse "No artists found"
// @Failure 500 {object} utils.ErrorResponse "Failed to load artists"
// @Router /api/v1/artists/all [get]
func (handlers *artistHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	artists, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get artists: %v", err))
		return
	} else if len(artists) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No artists were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(artists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode artists: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

// AddFavoriteArtist godoc
// @Summary Add favorite artist for user
// @Description Add new favorite artist for user.
// @Param artistID path int true "Artist ID"
// @Success 200
// @Failure 404 {object} utils.ErrorResponse "Invalid artist ID"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Can't add artist to favorite"
// @Router /api/v1/artists/favorite/{artistID} [post]
func (handlers *artistHandlers) AddFavoriteArtist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["artistID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid artist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.AddFavoriteArtist(request.Context(), userID, artistID); err != nil {
		handlers.logger.Error("Can't add artist to favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't add artist to favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// DeleteFavoriteArtist godoc
// @Summary Add favorite artist
// @Description Delete artist from favorite for user.
// @Param artistID path int true "Artist ID"
// @Success 200
// @Failure 404 {object} utils.ErrorResponse "Invalid artist ID"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Can't delete artist from favorite"
// @Router /api/v1/artists/favorite/{artistID} [delete]
func (handlers *artistHandlers) DeleteFavoriteArtist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["artistID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid artist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.DeleteFavoriteArtist(request.Context(), userID, artistID); err != nil {
		handlers.logger.Error("Can't delete artist from favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't delete artist from favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// IsFavoriteArtist godoc
// @Summary Check if an artist is a user's favorite
// @Description Checks if a specific artist is marked as a favorite for the authenticated user.
// @Param artistID path int true "Artist ID"
// @Success 200 {object} map[string]bool "Response indicating whether the artist is a favorite"
// @Failure 400 {object} utils.ErrorResponse "Invalid artist ID or user ID"
// @Failure 404 {object} utils.ErrorResponse "Artist ID not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/v1/artists/favorite/{artistID} [get]
func (handlers *artistHandlers) IsFavoriteArtist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["artistID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid artist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	exists, err := handlers.usecase.IsFavoriteArtist(request.Context(), userID, artistID)
	if err != nil {
		handlers.logger.Error("Can't check is artist in favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't check is artist in favorite")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(map[string]bool{"exists": exists}); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetFavoriteArtists godoc
// @Summary Get favorite artists
// @Description Retrieves a list of favorite artists for the user.
// @Success 200 {array} dto.ArtistDTO "List of favorite artists"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to get favorite artists"
// @Router /api/v1/artists/favorite [get]
func (handlers *artistHandlers) GetFavoriteArtists(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	artists, err := handlers.usecase.GetFavoriteArtists(request.Context(), userID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get favorite artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get favorite artists: %v", err))
		return
	} else if len(artists) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No favorite artists were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(artists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode artists: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}
