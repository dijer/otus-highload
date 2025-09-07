package handler_posts

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type getFeedRequest struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

func (h *PostsHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req getFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid request")
		return
	}

	posts, err := h.service.GetFeed(r.Context(), userID, req.Limit, req.Offset)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not get feed")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully get feed", posts)
}
