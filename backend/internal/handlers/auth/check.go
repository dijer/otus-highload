package handler_auth

import (
	"net/http"

	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

func (h *AuthHandler) CheckAuthHandler(w http.ResponseWriter, r *http.Request) {
	utils_server.JsonSuccess(w, http.StatusOK, "Success checked auth", nil)
}
