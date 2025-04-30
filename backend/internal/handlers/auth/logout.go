package handler_auth

import (
	"net/http"
	"time"

	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     h.authCfg.JWTCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	})

	utils_server.JsonSuccess(w, http.StatusOK, "Login successful", nil)
}
