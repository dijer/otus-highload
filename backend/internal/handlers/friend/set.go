package handler_friend

import (
	"encoding/json"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
)

type addFriendRequest struct {
	FriendID int `json:"friendId"`
}

func (h *FriendHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req addFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid friendId")
		return
	}

	if err := h.service.AddFriend(r.Context(), userID, req.FriendID); err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not add friend")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully added friend ", nil)
}
