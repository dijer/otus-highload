package handler_dialogs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/gorilla/mux"
)

type sendMessageRequest struct {
	Text string `json:"text"`
}

const userIDParam = "userId"

func (h *DialogsHandler) Send(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	toUserIDStr, ok := vars[userIDParam]
	if !ok {
		utils_server.JsonError(w, http.StatusBadRequest, "userId required")
		return
	}

	toUserID, err := strconv.ParseInt(toUserIDStr, 10, 64)
	if err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid userId")
		return
	}

	var req sendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils_server.JsonError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if req.Text == "" {
		utils_server.JsonError(w, http.StatusBadRequest, "text is required")
		return
	}

	err = h.service.Send(r.Context(), userID, toUserID, req.Text)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not create post")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully send message", nil)
}
