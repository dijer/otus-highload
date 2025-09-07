package service_friend

import (
	"context"

	storage_friend "github.com/dijer/otus-highload/backend/internal/storage/friend"
)

type FriendService struct {
	storage *storage_friend.FriendStorage
}

func New(storage *storage_friend.FriendStorage) *FriendService {
	return &FriendService{
		storage: storage,
	}
}

func (s *FriendService) AddFriend(ctx context.Context, userID, friendID int) error {
	return s.storage.AddFriend(ctx, userID, friendID)
}

func (s *FriendService) DeleteFriend(ctx context.Context, userID, friendID int) error {
	return s.storage.DeleteFriend(ctx, userID, friendID)
}
