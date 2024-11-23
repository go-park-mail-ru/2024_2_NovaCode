package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/csat"
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

func (handlers *csatHandlers) GetStatistics(response http.ResponseWriter, request *http.Request) {
	requestID := request.Context().Value(utils.RequestIDKey{})
	stat, err := handlers.usecase.GetStatistics(request.Context())
	if err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed get statistics: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Can't get statistics")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(stat); err != nil {
		handlers.logger.Error(fmt.Sprintf("Failed to encode statistics: %v", err), requestID)
		utils.JSONError(response, http.StatusInternalServerError, "Encode fail")
		return
	}

	response.WriteHeader(http.StatusOK)
}
