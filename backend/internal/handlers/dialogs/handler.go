package handler_dialogs

import (
	service_dialogs "github.com/dijer/otus-highload/backend/internal/services/dialogs"
)

type DialogsHandler struct {
	service *service_dialogs.DialogsService
}

func New(service *service_dialogs.DialogsService) *DialogsHandler {
	return &DialogsHandler{
		service: service,
	}
}
