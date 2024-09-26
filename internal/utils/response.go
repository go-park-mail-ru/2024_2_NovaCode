package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Error: message}
}

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) *MessageResponse {
	return &MessageResponse{message}
}

func JSONError(response http.ResponseWriter, statusCode int, message string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	errorResponse := NewErrorResponse(message)
	if err := json.NewEncoder(response).Encode(errorResponse); err != nil {
		http.Error(response, "failed to encode error message", http.StatusInternalServerError)
	}
}
