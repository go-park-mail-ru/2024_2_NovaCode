package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type trackHandlers struct {
	usecase track.Usecase
	logger  logger.Logger
}

func NewTrackHandlers(usecase track.Usecase, logger logger.Logger) track.Handlers {
	return &trackHandlers{usecase, logger}
}

// SearchTrack godoc
// @Summary Search tracks by name
// @Description Searches for tracks based on the provided "name" query parameter.
// @Param name query string true "Name of the track to search for"
// @Success 200 {array}  dto.TrackDTO "List of found tracks"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No tracks found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode tracks"
// @Router /api/v1/tracks/search [get]
func (handlers *trackHandlers) SearchTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	name := request.URL.Query().Get("name")
	if name == "" {
		handlers.logger.Error("Missing query parameter 'name'", requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundTracks, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't find tracks")
		return
	} else if len(foundTracks) == 0 {
		handlers.logger.Error(fmt.Sprintf("No tracks with %s were found", name), requestID)
		utils.JSONError(response, http.StatusNotFound, "No tracks")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundTracks); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// ViewTrack godoc
// @Summary Get track by ID
// @Description Retrieves an track using the provided track ID in the URL path.
// @Param id path uint64 true "Track ID"
// @Success 200 {object} dto.TrackDTO "Track found"
// @Failure 400 {object} utils.ErrorResponse "Invalid track ID"
// @Failure 404 {object} utils.ErrorResponse "Track not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to encode the track data"
// @Router /api/v1/tracks/{id} [get]
func (handlers *trackHandlers) ViewTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	trackID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundTrack, err := handlers.usecase.View(request.Context(), trackID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Track wasn't found: %v", err), requestID)
		utils.JSONError(response, http.StatusNotFound, "Can't find track")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundTrack); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode track: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetAll godoc
// @Summary Get all tracks
// @Description Retrieves a list of all tracks from the database.
// @Success 200 {array} dto.TrackDTO "List of all tracks"
// @Failure 404 {object} utils.ErrorResponse "No tracks found"
// @Failure 500 {object} utils.ErrorResponse "Failed to load tracks"
// @Router /api/v1/tracks/all [get]
func (handlers *trackHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	tracks, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get tracks: %v", err))
		return
	} else if len(tracks) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No tracks were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(tracks); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

// GetAllByArtistID godoc
// @Summary Get all tracks by artist ID
// @Description Retrieves a list of all tracks for a given artist ID.
// @Param artistId path int true "Artist ID"
// @Success 200 {array} dto.TrackDTO "List of tracks by artist"
// @Failure 404 {object} utils.ErrorResponse "No tracks found for the given artist ID"
// @Failure 500 {object} utils.ErrorResponse "Failed to load tracks by artist ID"
// @Router /api/v1/tracks/byArtistId/{artistId} [get]
func (handlers *trackHandlers) GetAllByArtistID(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistIDStr := vars["artistId"]
	artistID, err := strconv.ParseUint(artistIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid artist ID: %v", err))
		return
	}

	tracks, err := handlers.usecase.GetAllByArtistID(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get tracks by artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get tracks by artist ID: %v", err))
		return
	} else if len(tracks) == 0 {
		utils.JSONError(response, http.StatusNotFound, fmt.Sprintf("No tracks found for artist ID %d", artistID))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(tracks); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}
