package handler_auth

import (
	"net/http"
	"time"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	uuid := httpctx.GetUUID(r)
	if uuid == "" {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := h.cache.DeleteSession(r.Context(), uuid); err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "Failed to delete session")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.authCfg.JWTCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	})

	utils_server.JsonSuccess(w, http.StatusOK, "Logout successful", nil)
}
