package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type albumHandlers struct {
	cfg     *config.Config
	usecase album.Usecase
	logger  logger.Logger
}

func NewAlbumHandlers(cfg *config.Config, usecase album.Usecase, logger logger.Logger) album.Handlers {
	return &albumHandlers{cfg, usecase, logger}
}

func (handlers *albumHandlers) SearchAlbum(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		utils.JSONError(response, http.StatusBadRequest, "Missing query parameter 'name'")
		return
	}

	foundAlbums, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to find albums: %v", err))
		return
	} else if len(foundAlbums) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundAlbums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode albums: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *albumHandlers) ViewAlbum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["album"], 10, 64)
	if err != nil {
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Wrong album id: %v", err))
		return
	}

	foundAlbum, err := handlers.usecase.View(request.Context(), albumID)
	if err != nil {
		utils.JSONError(response, http.StatusNotFound, fmt.Sprintf("Albums  wasn't found: %v", err))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundAlbum); err != nil {
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode album: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *albumHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	albums, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to load albums: %v", err))
		return
	} else if len(albums) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(albums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode albums: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
}
