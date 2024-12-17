package http

import (
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"
	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/dto"
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
// @Summary Search tracks by query
// @Description Searches for tracks based on the provided "query" query parameter.
// @Param query query string true "Query of the track to search for"
// @Success 200 {array} dto.TrackDTO "List of found tracks"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No tracks found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode tracks"
// @Router /api/v1/tracks/search [get]
func (handlers *trackHandlers) SearchTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	query := request.URL.Query().Get("query")
	if query == "" {
		handlers.logger.Error("Missing query parameter 'query'", requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundTracks, err := handlers.usecase.Search(request.Context(), query)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't find tracks")
		return
	} else if len(foundTracks) == 0 {
		handlers.logger.Error(fmt.Sprintf("No tracks with %s were found", query), requestID)
		utils.JSONError(response, http.StatusNotFound, "No tracks")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(foundTracks))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
	response.Write(rawBytes)
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
	rawBytes, err := easyjson.Marshal(foundTrack)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode track: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
	response.Write(rawBytes)
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
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(tracks))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}
	response.Write(rawBytes)
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
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(tracks))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

// GetAllByAlbumID godoc
// @Summary Get all tracks by album ID
// @Description Retrieves a list of all tracks for a given album ID.
// @Param albumId path int true "Album ID"
// @Success 200 {array} dto.TrackDTO "List of tracks by album"
// @Failure 404 {object} utils.ErrorResponse "No tracks found for the given album ID"
// @Failure 500 {object} utils.ErrorResponse "Failed to load tracks by album ID"
// @Router /api/v1/tracks/byAlbumId/{albumId} [get]
func (handlers *trackHandlers) GetAllByAlbumID(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	albumIDStr := vars["albumId"]
	albumID, err := strconv.ParseUint(albumIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid album ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid album ID: %v", err))
		return
	}

	tracks, err := handlers.usecase.GetAllByAlbumID(request.Context(), albumID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get tracks by album ID: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get tracks by album ID: %v", err))
		return
	} else if len(tracks) == 0 {
		utils.JSONError(response, http.StatusNotFound, fmt.Sprintf("No tracks found for album ID %d", albumID))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(tracks))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

// AddFavoriteTrack godoc
// @Summary Add favorite track for user
// @Description Add new favorite track for user.
// @Param trackID path int true "Track ID"
// @Success 200
// @Failure 404 {object} utils.ErrorResponse "Invalid track ID"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Can't add track to favorite"
// @Router /api/v1/tracks/favorite/{trackID} [post]
func (handlers *trackHandlers) AddFavoriteTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	trackID, err := strconv.ParseUint(vars["trackID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid track ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid track ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.AddFavoriteTrack(request.Context(), userID, trackID); err != nil {
		handlers.logger.Error("Can't add track to favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't add track to favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// DeleteFavoriteTrack godoc
// @Summary Delete favorite track
// @Description Delete track from favorite for user.
// @Param trackID path int true "Track ID"
// @Success 200
// @Failure 404 {object} utils.ErrorResponse "Invalid track ID"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Can't delete track from favorite"
// @Router /api/v1/tracks/favorite/{trackID} [delete]
func (handlers *trackHandlers) DeleteFavoriteTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	trackID, err := strconv.ParseUint(vars["trackID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid track ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid track ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.DeleteFavoriteTrack(request.Context(), userID, trackID); err != nil {
		handlers.logger.Error("Can't delete track from favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't delete track from favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// IsFavoriteTrack godoc
// @Summary Check if a track is a user's favorite
// @Description Checks if a specific track is marked as a favorite for the authenticated user.
// @Param trackID path int true "Track ID"
// @Success 200 {object} map[string]bool "Response indicating whether the track is a favorite"
// @Failure 400 {object} utils.ErrorResponse "Invalid track ID or user ID"
// @Failure 404 {object} utils.ErrorResponse "Track ID not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/v1/tracks/favorite/{trackID} [get]
func (handlers *trackHandlers) IsFavoriteTrack(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	trackID, err := strconv.ParseUint(vars["trackID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid track ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid track ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	exists, err := handlers.usecase.IsFavoriteTrack(request.Context(), userID, trackID)
	if err != nil {
		handlers.logger.Error("Can't check is track in favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't check is track in favorite")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	existsResponse := &utils.ExistsResponse{Exists: exists}
	rawBytes, err := easyjson.Marshal(existsResponse)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode: %v", err))
		return
	}
	response.Write(rawBytes)

	response.WriteHeader(http.StatusOK)
}

// GetFavoriteTracks godoc
// @Summary Get favorite tracks
// @Description Retrieves a list of favorite tracks for the user.
// @Success 200 {array} dto.TrackDTO "List of favorite tracks"
// @Failure 404 {object} utils.ErrorResponse "User id not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to get favorite tracks"
// @Router /api/v1/tracks/favorite [get]
func (handlers *trackHandlers) GetFavoriteTracks(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	tracks, err := handlers.usecase.GetFavoriteTracks(request.Context(), userID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get favorite tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get favorite tracks: %v", err))
		return
	} else if len(tracks) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No favorite tracks were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(tracks))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *trackHandlers) GetTracksFromPlaylist(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playlistIDStr := vars["playlistId"]
	playlistID, err := strconv.ParseUint(playlistIDStr, 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	playlist, err := h.usecase.GetTracksFromPlaylist(request.Context(), playlistID)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.TrackDTOs(playlist))
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}
