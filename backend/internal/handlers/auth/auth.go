package handler_auth

import (
	"github.com/dijer/otus-highload/backend/internal/config"
	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
)

type AuthHandler struct {
	service *service_user.UserService
	authCfg config.AuthConf
}

func New(service *service_user.UserService, authCfg config.AuthConf) *AuthHandler {
	return &AuthHandler{
		service: service,
		authCfg: authCfg,
	}
}
