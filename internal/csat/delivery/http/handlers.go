package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type csatHandlers struct {
	usecase csat.Usecase
	logger  logger.Logger
}

func NewCSATHandlers(usecase csat.Usecase, logger logger.Logger) csat.Handlers {
	return &csatHandlers{usecase, logger}
}

func (handlers *csatHandlers) SubmitAnswer(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})

	userID, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}

	vars := mux.Vars(request)
	questionID, err := strconv.ParseUint(vars["questionID"], 10, 64)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("get '%s' wrong id: %v", vars["questionID"], err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "wrong id value")
		return
	}

	var csatAnswerDTO dto.CSATAnswerDTO
	if err := json.NewDecoder(request.Body).Decode(&csatAnswerDTO); err != nil {
		handlers.logger.Error(fmt.Sprintf("invalid request body: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}
	csatAnswerDTO.CSATQuestionID = questionID
	csatAnswerDTO.UserID = userID

	answer := dto.NewAnswerFromCSATAnswerDTO(&csatAnswerDTO)
	answerDTO, err := handlers.usecase.SubmitAnswer(request.Context(), answer)
	if err != nil {
		handlers.logger.Error("cannot submit answer for question", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user has already answered")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(answerDTO); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return updated user details")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handlers *csatHandlers) GetQuestionsByTopic(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})

	_, ok := request.Context().Value(utils.UserIDKey{}).(uuid.UUID)
	if !ok {
		handlers.logger.Error("user id not found in context", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user id not found")
		return
	}

	topic := request.URL.Query().Get("topic")
	if topic == "" {
		handlers.logger.Error("missing query parameter 'topic'", requestID)
		utils.JSONError(response, http.StatusBadRequest, "invalid query")
		return
	}

	questions, err := handlers.usecase.GetQuestionsByTopic(request.Context(), topic)
	if err != nil {
		handlers.logger.Warn(fmt.Sprintf("failed to get csat questions: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to get csat questions")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(questions); err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return updated user details")
		return
	}

	response.WriteHeader(http.StatusOK)
}
