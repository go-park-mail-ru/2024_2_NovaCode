package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/csat/dto"
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

	csatAnswerDTO := &dto.CSATAnswerDTO{}
	rawBytes, _ := io.ReadAll(request.Body)
	err = easyjson.Unmarshal(rawBytes, csatAnswerDTO)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("invalid request body: %v", err), requestID)
		utils.JSONError(response, http.StatusBadRequest, "invalid request body")
		return
	}
	csatAnswerDTO.CSATQuestionID = questionID
	csatAnswerDTO.UserID = userID

	answer := dto.NewAnswerFromCSATAnswerDTO(csatAnswerDTO)
	answerDTO, err := handlers.usecase.SubmitAnswer(request.Context(), answer)
	if err != nil {
		handlers.logger.Error("cannot submit answer for question", requestID)
		utils.JSONError(response, http.StatusBadRequest, "user has already answered")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err = easyjson.Marshal(answerDTO)
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return updated user details")
		return
	}
	response.Write(rawBytes)
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
	rawBytes, err := easyjson.Marshal(dto.CSATQuestionDTOs(questions))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("error encoding updated user response: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "failed to return updated user details")
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}

func (handlers *csatHandlers) GetStatistics(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	stat, err := handlers.usecase.GetStatistics(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed get statistics: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't get statistics")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	rawBytes, err := easyjson.Marshal(dto.CSATStatisticsDTOs(stat))
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode statistics: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}
	response.Write(rawBytes)
	response.WriteHeader(http.StatusOK)
}
