package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
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
	requestId := uuid.New()
	ctx := context.WithValue(request.Context(), utils.RequestIdKey{}, requestId)

	name := request.URL.Query().Get("name")
	if name == "" {
		handlers.logger.Error("Missing query parameter 'name'", zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundAlbums, err := handlers.usecase.Search(ctx, name)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find albums: %v", err), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusInternalServerError, "Can't find albums")
		return
	} else if len(foundAlbums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums with %s were found", name), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusNotFound, "No albums")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, foundAlbums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), zap.String("request_id", requestId.String()))
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
	requestId := uuid.New()
	ctx := context.WithValue(request.Context(), utils.RequestIdKey{}, requestId)

	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundAlbum, err := handlers.usecase.View(ctx, albumID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Album wasn't found: %v", err), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusNotFound, "Can't find album")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, foundAlbum); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode album: %v", err), zap.String("request_id", requestId.String()))
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
	requestId := uuid.New()
	ctx := context.WithValue(request.Context(), utils.RequestIdKey{}, requestId)

	albums, err := handlers.usecase.GetAll(ctx)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums: %v", err), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusInternalServerError, "Albums load fail")
		return
	} else if len(albums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums were found"), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusNotFound, "No albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := utils.WriteResponse(response, http.StatusOK, albums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), zap.String("request_id", requestId.String()))
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
}

// GetAllByArtistID godoc
// @Summary Get all albums by artist ID
// @Description Retrieves a list of all albums for a given artist ID from the database.
// @Param artistID path int true "Artist ID"
// @Success 200 {array} dto.AlbumDTO "List of albums by artist"
// @Failure 404 {object} utils.ErrorResponse "No albums found for the artist"
// @Failure 500 {object} utils.ErrorResponse "Failed to load albums"
// @Router /api/v1/albums/artist/{artistID} [get]
func (handlers *albumHandlers) GetAllByArtistID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	artistIDStr := vars["artistId"]
	artistID, err := strconv.ParseUint(artistIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err))
		utils.JSONError(response, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	albums, err := handlers.usecase.GetAllByArtistID(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums by artist ID: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Albums load fail")
		return
	} else if len(albums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums found for artist ID: %d", artistID))
		utils.JSONError(response, http.StatusNotFound, "No albums found for the artist")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(albums); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err))
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}
