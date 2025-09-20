package handler_posts

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type createPostRequest struct {
	Content string `json:"content"`
}

func (h *PostsHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req createPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid content")
		return
	}

	post, err := h.service.CreatePost(r.Context(), userID, req.Content)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not create post")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully created post", post.ID)
}
