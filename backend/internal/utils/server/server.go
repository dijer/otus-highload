package utils_server

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/models"
)

func JsonError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := models.ServerResponse{
		Ok:      false,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func JsonSuccess(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := models.ServerResponse{
		Ok:      true,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
