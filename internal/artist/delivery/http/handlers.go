package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
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
// @Summary Search artists by name
// @Description Searches for artists based on the provided "name" query parameter.
// @Param name query string true "Name of the artist to search for"
// @Success 200 {array}  dto.ArtistDTO "List of found artists"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No artists found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode artists"
// @Router /api/v1/artists/search [get]
func (handlers *artistHandlers) SearchArtist(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		handlers.logger.Error("Missing query parameter 'name'")
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundArtists, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find artists: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Can't find artists")
		return
	} else if len(foundArtists) == 0 {
		handlers.logger.Error(fmt.Sprintf("No artists with %s were found", name))
		utils.JSONError(response, http.StatusNotFound, "No artists")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err))
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
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err))
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundArtist, err := handlers.usecase.View(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Arist wasn't found: %v", err))
		utils.JSONError(response, http.StatusNotFound, "Can't find artist")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtist); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artist: %v", err))
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
	artists, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get artists: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get artists: %v", err))
		return
	} else if len(artists) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No artists were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(artists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode artists: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}
