package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

	if err := json.NewDecoder(request.Body).Decode(playlistDTO); err != nil {
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
	if err = json.NewEncoder(response).Encode(newPlaylist); err != nil {
		h.logger.Errorf("can't encode")
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (h *playlistHandlers) GetAllPlaylists(response http.ResponseWriter, request *http.Request) {
	playlists, err := h.usecase.GetAllPlaylists(request.Context())
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(response).Encode(playlists); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

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
	if err = json.NewEncoder(response).Encode(playlist); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

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
	if err = json.NewEncoder(response).Encode(playlists); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

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

	payload := &struct {
		TrackID uint64 `json:"track_id"`
	}{}

	if err := json.NewDecoder(request.Body).Decode(payload); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	playlistTrackDTO := &dto.PlaylistTrackDTO{PlaylistID: playlistID, TrackID: payload.TrackID}

	playlistTrack, err := h.usecase.AddToPlaylist(request.Context(), playlistTrackDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(response).Encode(playlistTrack); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

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

	payload := &struct {
		TrackID uint64 `json:"track_id"`
	}{}

	if err := json.NewDecoder(request.Body).Decode(payload); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	playlistTrackDTO := &dto.PlaylistTrackDTO{PlaylistID: playlistID, TrackID: payload.TrackID}

	err = h.usecase.RemoveFromPlaylist(request.Context(), playlistTrackDTO)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json")
	message := utils.NewMessageResponse("")
	if err := json.NewEncoder(response).Encode(message); err != nil {
		utils.JSONError(response, http.StatusBadRequest, "")
		return
	}

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
	message := utils.NewMessageResponse("")
	if err = json.NewEncoder(response).Encode(message); err != nil {
		utils.JSONError(response, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
}
