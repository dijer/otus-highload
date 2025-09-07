package handler_posts

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type getPostRequest struct {
	PostID int `json:"postId"`
}

func (h *PostsHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req getPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid postId")
		return
	}

	post, err := h.service.GetPost(r.Context(), userID, req.PostID)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not get post")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully get post", post)
}
