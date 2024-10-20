package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
)

type albumHandlers struct {
	usecase album.Usecase
	logger  logger.Logger
}

func NewAlbumHandlers(usecase album.Usecase, logger logger.Logger) album.Handlers {
	return &albumHandlers{usecase, logger}
}

// SearchAlbum godoc
// @Summary Search albums by name
// @Description Searches for albums based on the provided "name" query parameter.
// @Param name query string true "Name of the album to search for"
// @Success 200 {array}  dto.AlbumDTO "List of found albums"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No albums found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode albums"
// @Router /api/v1/albums/search [get]
func (handlers *albumHandlers) SearchAlbum(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		handlers.logger.Error("Missing query parameter 'name'")
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundAlbums, err := handlers.usecase.Search(request.Context(), name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Can't find albums")
		return
	} else if len(foundAlbums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums with %s were found", name))
		utils.JSONError(response, http.StatusNotFound, "No albums")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, foundAlbums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}

// ViewAlbum godoc
// @Summary Get album by ID
// @Description Retrieves an album using the provided album ID in the URL path.
// @Param id path uint64 true "Album ID"
// @Success 200 {object} dto.AlbumDTO "Album found"
// @Failure 400 {object} utils.ErrorResponse "Invalid album ID"
// @Failure 404 {object} utils.ErrorResponse "Album not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to encode the album data"
// @Router /api/v1/albums/{id} [get]
func (handlers *albumHandlers) ViewAlbum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err))
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundAlbum, err := handlers.usecase.View(request.Context(), albumID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Album wasn't found: %v", err))
		utils.JSONError(response, http.StatusNotFound, "Can't find album")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, foundAlbum); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode album: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
}

// GetAll godoc
// @Summary Get all albums
// @Description Retrieves a list of all albums from the database.
// @Success 200 {array} dto.AlbumDTO "List of all albums"
// @Failure 404 {object} utils.ErrorResponse "No albums found"
// @Failure 500 {object} utils.ErrorResponse "Failed to load albums"
// @Router /api/v1/albums/all [get]
func (handlers *albumHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	albums, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Albums load fail")
		return
	} else if len(albums) == 0 {
		handlers.logger.Error("No albums were found")
		utils.JSONError(response, http.StatusNotFound, "No albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, albums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
}
