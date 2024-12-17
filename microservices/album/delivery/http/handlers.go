package http

import (
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type albumHandlers struct {
	usecase album.Usecase
	logger  logger.Logger
}

func NewAlbumHandlers(usecase album.Usecase, logger logger.Logger) album.Handlers {
	return &albumHandlers{usecase, logger}
}

// SearchAlbum godoc
// @Summary Search albums by query
// @Description Searches for albums based on the provided "query" query parameter.
// @Param query query string true "Name of the album to search for"
// @Success 200 {array}  dto.AlbumDTO "List of found albums"
// @Failure 400 {object} utils.ErrorResponse "Missing or invalid query parameter"
// @Failure 404 {object} utils.ErrorResponse "No albums found with the provided name"
// @Failure 500 {object} utils.ErrorResponse "Failed to search or encode albums"
// @Router /api/v1/albums/search [get]
func (handlers *albumHandlers) SearchAlbum(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	query := request.URL.Query().Get("query")
	if query == "" {
		handlers.logger.Error("Missing query parameter 'query'", requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong query")
		return
	}

	foundAlbums, err := handlers.usecase.Search(request.Context(), query)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to find albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't find albums")
		return
	} else if len(foundAlbums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums with %s were found", query), requestID)
		utils.JSONError(response, http.StatusNotFound, "No albums")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.AlbumDTOs(foundAlbums))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
		return
	}
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
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Get '%s' wrong id: %v", vars["id"], err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Wrong id value")
		return
	}

	foundAlbum, err := handlers.usecase.View(request.Context(), albumID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Album wasn't found: %v", err), requestID)
		utils.JSONError(response, http.StatusNotFound, "Can't find album")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(foundAlbum)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode album: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
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
	requestID := request.Context().Value(utils.RequestIDKey{})
	albums, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Albums load fail")
		return
	} else if len(albums) == 0 {
		handlers.logger.Error("No albums were found", requestID)
		utils.JSONError(response, http.StatusNotFound, "No albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.AlbumDTOs(albums))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
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
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistIDStr := vars["artistId"]
	artistID, err := strconv.ParseUint(artistIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	albums, err := handlers.usecase.GetAllByArtistID(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load albums by artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Albums load fail")
		return
	} else if len(albums) == 0 {
		handlers.logger.Error(fmt.Sprintf("No albums found for artist ID: %d", artistID), requestID)
		utils.JSONError(response, http.StatusNotFound, "No albums found for the artist")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.AlbumDTOs(albums))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
		return
	}
}

func (handlers *albumHandlers) AddFavoriteAlbum(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["albumID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid album ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid album ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.AddFavoriteAlbum(request.Context(), userID, albumID); err != nil {
		handlers.logger.Error("Can't add album to favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't add album to favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *albumHandlers) DeleteFavoriteAlbum(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["albumID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid album ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid album ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	if err := handlers.usecase.DeleteFavoriteAlbum(request.Context(), userID, albumID); err != nil {
		handlers.logger.Error("Can't delete album from favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't delete album from favorite")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *albumHandlers) IsFavoriteAlbum(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	albumID, err := strconv.ParseUint(vars["albumID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid album ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, fmt.Sprintf("Invalid album ID: %v", err))
		return
	}

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	exists, err := handlers.usecase.IsFavoriteAlbum(request.Context(), userID, albumID)
	if err != nil {
		handlers.logger.Error("Can't check is album in favorite", requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't check is album in favorite")
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

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
		return
	}
}

func (handlers *albumHandlers) GetFavoriteAlbums(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("User id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "User id not found")
		return
	}

	albums, err := handlers.usecase.GetFavoriteAlbums(request.Context(), userID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to get favorite albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to get favorite albums: %v", err))
		return
	} else if len(albums) == 0 {
		utils.JSONError(response, http.StatusNotFound, "No favorite albums were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.AlbumDTOs(albums))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode albums: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, fmt.Sprintf("Failed to encode albums: %v", err))
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(rawBytes)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to write response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Write response fail")
		return
	}
}
