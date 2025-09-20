package handler_friend

import (
	service_friend "github.com/dijer/otus-highload/backend/internal/services/friend"
)

type FriendHandler struct {
	service *service_friend.FriendService
}

func New(service *service_friend.FriendService) *FriendHandler {
	return &FriendHandler{
		service: service,
	}
}
