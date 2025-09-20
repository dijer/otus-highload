package handler_posts

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type updatePostRequest struct {
	PostID  int64  `json:"postId"`
	Content string `json:"content"`
}

func (h *PostsHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req updatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := h.service.UpdatePost(r.Context(), userID, req.PostID, req.Content); err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not update post")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully updated post", nil)
}
