package handler_auth

import (
	cache_auth "github.com/dijer/otus-highload/backend/internal/cache/auth"
	"github.com/dijer/otus-highload/backend/internal/config"
	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
)

type AuthHandler struct {
	service *service_user.UserService
	authCfg config.AuthConf
	cache   *cache_auth.AuthCache
}

func New(service *service_user.UserService, authCfg config.AuthConf, cache *cache_auth.AuthCache) *AuthHandler {
	return &AuthHandler{
		service: service,
		authCfg: authCfg,
		cache:   cache,
	}
}
