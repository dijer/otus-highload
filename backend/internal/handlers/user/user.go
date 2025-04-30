package handler_user

import (
	"net/http"
	"strconv"

	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *service_user.UserService
}

func New(service *service_user.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Invalid userID")
		return
	}

	user, err := h.service.GetUser(userID)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Cant get user by userID")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Get user successfully", user)
}
