package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type artistHandlers struct {
	cfg     *config.Config
	usecase artist.Usecase
	logger  logger.Logger
}

func NewArtistHandlers(cfg *config.Config, usecase artist.Usecase, logger logger.Logger) artist.Handlers {
	return &artistHandlers{cfg, usecase, logger}
}

func (handlers *artistHandlers) SearchArtist(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		utils.JSONError(response, http.StatusBadRequest, "Missing query parameter 'name'")
		return
	}

	foundArtists, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find artists: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to find artists: %v", err))
		return
	} else if len(foundArtists) == 0 {
		utils.JSONError(response, http.StatusNotFound, "")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtists); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode artists: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode artists: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *artistHandlers) ViewArtist(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	artistID, err := strconv.ParseUint(vars["artist"], 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Wrong artist id: %v", err))
		return
	}

	foundArtist, err := handlers.usecase.View(request.Context(), artistID)
	if err != nil {
		utils.JSONError(response, http.StatusNotFound, fmt.Sprintf("Artist wasn't found: %v", err))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundArtist); err != nil {
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode artist: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}
