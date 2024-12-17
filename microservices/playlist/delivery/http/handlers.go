package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type playlistHandlers struct {
	usecase playlist.Usecase
	logger  logger.Logger
}

func NewPlaylistHandlers(usecase playlist.Usecase, logger logger.Logger) playlist.Handlers {
	return &playlistHandlers{usecase: usecase, logger: logger}
}

func (h *playlistHandlers) CreatePlaylist(response http.ResponseWriter, request *http.Request) {
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		h.logger.Errorf("no user, userID: %v", userID)
		utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
		return
	}

	playlistDTO := &dto.PlaylistDTO{}
	rawBytes, _ := io.ReadAll(request.Body)
	err := easyjson.Unmarshal(rawBytes, playlistDTO)
	if err != nil {
		h.logger.Errorf("can't decode")
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	playlistDTO.OwnerID = userID
	newPlaylist, err := h.usecase.CreatePlaylist(request.Context(), playlistDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err = easyjson.Marshal(newPlaylist)
	if err != nil {
		h.logger.Errorf("can't encode")
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) GetAllPlaylists(response http.ResponseWriter, request *http.Request) {
	playlists, err := h.usecase.GetAllPlaylists(request.Context())
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.PlaylistDTOs(playlists))
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) GetPlaylist(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playlistIDStr := vars["playlistId"]
	playlistID, err := strconv.ParseUint(playlistIDStr, 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	playlist, err := h.usecase.GetPlaylist(request.Context(), playlistID)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(playlist)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) GetUserPlaylists(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userIDStr := vars["userId"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	playlists, err := h.usecase.GetUserPlaylists(request.Context(), userID)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.PlaylistDTOs(playlists))
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) AddToPlaylist(response http.ResponseWriter, request *http.Request) {
	_, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
		return
	}

	vars := mux.Vars(request)
	playlistIDStr := vars["playlistId"]
	playlistID, err := strconv.ParseUint(playlistIDStr, 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	trackIdDTO := &dto.TrackIdDTO{}

	rawBytes, _ := io.ReadAll(request.Body)
	err = easyjson.Unmarshal(rawBytes, trackIdDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	playlistTrackDTO := &dto.PlaylistTrackDTO{PlaylistID: playlistID, TrackID: trackIdDTO.TrackID}

	playlistTrack, err := h.usecase.AddToPlaylist(request.Context(), playlistTrackDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err = easyjson.Marshal(playlistTrack)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) RemoveFromPlaylist(response http.ResponseWriter, request *http.Request) {
	_, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
		return
	}

	vars := mux.Vars(request)
	playlistIDStr := vars["playlistId"]
	playlistID, err := strconv.ParseUint(playlistIDStr, 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	trackIdDTO := &dto.TrackIdDTO{}
	rawBytes, _ := io.ReadAll(request.Body)
	err = easyjson.Unmarshal(rawBytes, trackIdDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	playlistTrackDTO := &dto.PlaylistTrackDTO{PlaylistID: playlistID, TrackID: trackIdDTO.TrackID}

	err = h.usecase.RemoveFromPlaylist(request.Context(), playlistTrackDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	message := &utils.MessageResponse{}
	rawBytes, err = easyjson.Marshal(message)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "")
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) DeletePlaylist(response http.ResponseWriter, request *http.Request) {
	_, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		utils.JSONError(response, http.StatusUnauthorized, "unauthorized")
		return
	}

	vars := mux.Vars(request)
	playlistIDStr := vars["playlistId"]
	playlistID, err := strconv.ParseUint(playlistIDStr, 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, "Invalid playlist ID")
		return
	}

	err = h.usecase.DeletePlaylist(request.Context(), playlistID)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	message := &utils.MessageResponse{}
	rawBytes, err := easyjson.Marshal(message)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) AddFavoritePlaylist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	playlistID, err := strconv.ParseUint(vars["playlistID"], 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Invalid playlist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid playlist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		h.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := h.usecase.AddFavoritePlaylist(request.Context(), userID, playlistID); err != nil {
		h.logger.Error("Can't add playlist to favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't add playlist to favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) DeleteFavoritePlaylist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	playlistID, err := strconv.ParseUint(vars["playlistID"], 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Invalid playlist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid playlist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		h.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := h.usecase.DeleteFavoritePlaylist(request.Context(), userID, playlistID); err != nil {
		h.logger.Error("Can't delete playlist from favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't delete playlist from favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) IsFavoritePlaylist(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	playlistID, err := strconv.ParseUint(vars["playlistID"], 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Invalid playlist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid playlist ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		h.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	exists, err := h.usecase.IsFavoritePlaylist(request.Context(), userID, playlistID)
	if err != nil {
		h.logger.Error("Can't check is playlist in favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't check is playlist in favorite")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	existsResponse := &utils.ExistsResponse{Exists: exists}
	rawBytes, err := easyjson.Marshal(existsResponse)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Failed to encode: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode: %v", err))
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) GetFavoritePlaylists(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		h.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	playlists, err := h.usecase.GetFavoritePlaylists(request.Context(), userID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Failed to get favorite playlists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get favorite playlists: %v", err))
		return
	} else if len(playlists) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No favorite playlists were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.PlaylistDTOs(playlists))
	if err != nil {
		h.logger.Error(fmt.Sprintf("Failed to encode playlists: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode playlists: %v", err))
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}
