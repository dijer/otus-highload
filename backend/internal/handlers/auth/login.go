package handler_auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dijer/otus-highload/backend/internal/models"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/golang-jwt/jwt/v5"
)

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserWithPassword
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Bad request")
		return
	}

	userID, err := h.service.CheckUserPassword(r.Context(), user)
	if err != nil {
		utils_server.JsonError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	tokenStr, err := h.generateJWT(userID)
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

	utils_server.JsonSuccess(w, http.StatusOK, "Login successful", nil)
}

func (h *AuthHandler) generateJWT(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Duration(h.authCfg.JWTExpireHours) * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(h.authCfg.JWTKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
