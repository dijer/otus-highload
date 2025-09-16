package storage_friend

import (
	"context"

	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
)

type FriendStorage struct {
	dbRouter infra_database.DBRouter
}

func New(dbRouter infra_database.DBRouter) *FriendStorage {
	return &FriendStorage{
		dbRouter: dbRouter,
	}
}

func (s *FriendStorage) AddFriend(ctx context.Context, userID, friendID int) error {
	_, err := s.dbRouter.Exec(ctx,
		`INSERT INTO follows (user_id, friend_id) 
			VALUES ($1,$2) ON CONFLICT DO NOTHING`,
		userID, friendID,
	)

	return err
}

func (s *FriendStorage) DeleteFriend(ctx context.Context, userID, friendID int) error {
	_, err := s.dbRouter.Exec(ctx,
		`DELETE FROM follows WHERE user_id = $1 AND friend_id = $2`,
		userID, friendID,
	)

	return err
}
