package handler_profile

import (
	"net/http"

	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
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
	userID := r.Context().Value("userId").(int)

	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Cant get user by userID")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Get user successfully", user)
}
