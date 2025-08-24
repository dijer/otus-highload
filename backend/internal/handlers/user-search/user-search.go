package user_search

import (
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/models"
	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"

	"github.com/gorilla/schema"
)

type UserSearchParams struct {
	FirstName string `schema:"firstname"`
	LastName  string `schema:"lastname"`
}

type UserSearchHandler struct {
	service *service_user.UserService
}

type UserSearchResponse struct {
	Count int           `json:"count"`
	Users []models.User `json:"users"`
}

func New(service *service_user.UserService) *UserSearchHandler {
	return &UserSearchHandler{
		service: service,
	}
}

func (h *UserSearchHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var params UserSearchParams
	decoder := schema.NewDecoder()
	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "Invalid search params")
		return
	}

	if params.FirstName == "" || params.LastName == "" {
		utils_server.JsonError(w, http.StatusBadRequest, "Empty firstname or lastname")
		return
	}

	users, err := h.service.GetUsers(r.Context(), params.FirstName, params.LastName)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "Failed get users by firstname and lastname")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Get users successfully", UserSearchResponse{
		Count: len(users),
		Users: users,
	})
}
