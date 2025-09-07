package handler_profile

import (
	"net/http"

	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type ProfileHandler struct {
	service *service_user.UserService
}

func New(service *service_user.UserService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func (h *ProfileHandler) Handler(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Cant get user by userID")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Get user successfully", user)
}
