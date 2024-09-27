package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/track"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type trackHandlers struct {
	cfg     *config.Config
	usecase track.Usecase
	logger  logger.Logger
}

func NewTrackHandlers(cfg *config.Config, usecase track.Usecase, logger logger.Logger) track.Handlers {
	return &trackHandlers{cfg, usecase, logger}
}

func (handlers *trackHandlers) SearchTrack(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		utils.JSONError(response, http.StatusBadRequest, "Missing query parameter 'name'")
		return
	}

	foundTracks, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find tracks: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to find tracks: %v", err))
		return
	} else if len(foundTracks) == 0 {
		utils.JSONError(response, http.StatusNotFound, "")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundTracks); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode tracks: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode tracks: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *trackHandlers) ViewTrack(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	trackID, err := strconv.ParseUint(vars["track"], 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Wrong track id: %v", err))
		return
	}

	foundTrack, err := handlers.usecase.View(request.Context(), trackID)
	if err != nil {
		utils.JSONError(response, http.StatusNotFound, fmt.Sprintf("Track wasn't found: %v", err))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundTrack); err != nil {
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode track: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}