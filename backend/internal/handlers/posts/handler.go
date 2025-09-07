package handler_posts

import (
	service_posts "github.com/dijer/otus-highload/backend/internal/services/posts"
)

type PostsHandler struct {
	service *service_posts.PostsService
}

func New(service *service_posts.PostsService) *PostsHandler {
	return &PostsHandler{
		service: service,
	}
}
