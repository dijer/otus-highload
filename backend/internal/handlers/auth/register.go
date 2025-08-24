package handler_auth

import (
	"encoding/json"
	"errors"
	"net/http"

	errs "github.com/dijer/otus-highload/backend/internal/errors"
	models "github.com/dijer/otus-highload/backend/internal/models"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserWithPassword
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateUser(r.Context(), user); err != nil {
		if errors.Is(err, errs.ErrUserAlreadyExists) {
			utils_server.JsonError(w, http.StatusConflict, "User already exists")
			return
		}

		utils_server.JsonError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils_server.JsonSuccess(w, http.StatusCreated, "User registered successfully", nil)
}
