package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/genre"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/genre/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type genreHandlers struct {
	usecase genre.Usecase
	logger  logger.Logger
}

func NewGenreHandlers(usecase genre.Usecase, logger logger.Logger) genre.Handlers {
	return &genreHandlers{usecase, logger}
}

// GetAll godoc
// @Summary Get all genres
// @Description Retrieves a list of all genres from the database.
// @Success 200 {array} dto.GenreDTO "List of all genres"
// @Failure 404 {object} utils.ErrorResponse "No genres found"
// @Failure 500 {object} utils.ErrorResponse "Failed to load genres"
// @Router /api/v1/genres/all [get]
func (handlers *genreHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	genres, err := handlers.usecase.GetAll(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load genres: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Genres load fail")
		return
	} else if len(genres) == 0 {
		handlers.logger.Error("No genres were found", requestID)
		utils.JSONError(response, http.StatusNotFound, "No genres were found")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.GenreDTOs(genres))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode genres: %v", err), requestID)
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
// @Summary Get all genres by artist ID
// @Description Retrieves a list of all genres for a given artist ID from the database.
// @Param artistID path int true "Artist ID"
// @Success 200 {array} dto.GenreDTO "List of genres by artist"
// @Failure 404 {object} utils.ErrorResponse "No genres found for the artist"
// @Failure 500 {object} utils.ErrorResponse "Failed to load genres"
// @Router /api/v1/genres/artist/{artistID} [get]
func (handlers *genreHandlers) GetAllByArtistID(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	artistIDStr := vars["artistId"]
	artistID, err := strconv.ParseUint(artistIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	genres, err := handlers.usecase.GetAllByArtistID(request.Context(), artistID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load genres by artist ID: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Genres load fail")
		return
	} else if len(genres) == 0 {
		handlers.logger.Error(fmt.Sprintf("No genres found for artist ID: %d", artistID), requestID)
		utils.JSONError(response, http.StatusNotFound, "No genres found for the artist")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.GenreDTOs(genres))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode genres: %v", err), requestID)
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

// GetAllByTrackID godoc
// @Summary Get all genres by track ID
// @Description Retrieves a list of all genres for a given track ID from the database.
// @Param trackID path int true "Track ID"
// @Success 200 {array} dto.GenreDTO "List of genres by track"
// @Failure 404 {object} utils.ErrorResponse "No genres found for the track"
// @Failure 500 {object} utils.ErrorResponse "Failed to load genres"
// @Router /api/v1/genres/track/{trackID} [get]
func (handlers *genreHandlers) GetAllByTrackID(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(request)
	trackIDStr := vars["trackId"]
	trackID, err := strconv.ParseUint(trackIDStr, 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Invalid track ID: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "Invalid track ID")
		return
	}

	genres, err := handlers.usecase.GetAllByTrackID(request.Context(), trackID)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to load genres by track ID: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Genres load fail")
		return
	} else if len(genres) == 0 {
		handlers.logger.Error(fmt.Sprintf("No genres found for track ID: %d", trackID), requestID)
		utils.JSONError(response, http.StatusNotFound, "No genres found for the track")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.GenreDTOs(genres))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode genres: %v", err), requestID)
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
