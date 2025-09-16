package handler_posts

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type deletePostRequest struct {
	PostID int64 `json:"postId"`
}

func (h *PostsHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req deletePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid postId")
		return
	}

	if err := h.service.DeletePost(r.Context(), userID, req.PostID); err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not delete post")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully deleted post", nil)
}
