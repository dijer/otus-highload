package handler_dialogs

import (
	"net/http"
	"strconv"

	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/gorilla/mux"
)

func (h *DialogsHandler) List(w http.ResponseWriter, r *http.Request) {
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

	messages, err := h.service.List(r.Context(), userID, toUserID)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "can not get messages")
		return
	}

	utils_server.JsonSuccess(w, http.StatusOK, "Successfully get messages", messages)
}
