package utils

import (
	"encoding/json"
	"net/http"
)

func JSONError(response http.ResponseWriter, statusCode int, message string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(map[string]string{"error": message})
}
