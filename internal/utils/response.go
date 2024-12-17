package utils

import (
	"net/http"

	"github.com/mailru/easyjson"
)

//easyjson:json
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Error: message}
}

//easyjson:json
type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) *MessageResponse {
	return &MessageResponse{message}
}

//easyjson:json
type CSRFResponse struct {
	CSRF string `json:"csrf"`
}

func NewCSRFResponse(csrf string) *CSRFResponse {
	return &CSRFResponse{csrf}
}

func JSONError(response http.ResponseWriter, statusCode int, message string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	errorResponse := NewErrorResponse(message)
	rawBytes, err := easyjson.Marshal(errorResponse)
	if err != nil {
		http.Error(response, "failed to encode error message", http.StatusInternalServerError)
	}
	_, _ = response.Write(rawBytes)

}

//easyjson:json
type ExistsResponse struct {
	Exists bool `json:"exists"`
}
