package handler_auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	errs "github.com/dijer/otus-highload/backend/internal/errors"
	models "github.com/dijer/otus-highload/backend/internal/models"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/google/uuid"
)

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserWithPassword
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	userID, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, errs.ErrUserAlreadyExists) {
			utils_server.JsonError(w, http.StatusConflict, "User already exists")
			return
		}

		utils_server.JsonError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	sessionUUID := uuid.NewString()
	ttl := time.Hour * time.Duration(h.authCfg.JWTExpireHours)
	if err := h.cache.SaveSession(r.Context(), sessionUUID, userID, ttl); err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "cache save session error")
		return
	}

	tokenStr, err := h.generateJWT(userID, sessionUUID)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.authCfg.JWTCookieName,
		Value:    tokenStr,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * time.Duration(h.authCfg.JWTExpireHours)),
		HttpOnly: true,
	})

	utils_server.JsonSuccess(w, http.StatusCreated, "User registered successfully", nil)
}
