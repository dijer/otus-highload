package service_dialogs

import (
	"context"

	"github.com/dijer/otus-highload/backend/internal/models"
	storage_dialogs "github.com/dijer/otus-highload/backend/internal/storage/dialogs"
)

type DialogsService struct {
	storage *storage_dialogs.DialogsStorage
}

func New(storage *storage_dialogs.DialogsStorage) *DialogsService {
	return &DialogsService{
		storage: storage,
	}
}

func (s *DialogsService) Send(ctx context.Context, user1, user2 int64, text string) error {
	return s.storage.Send(ctx, user1, user2, text)
}

func (s *DialogsService) List(ctx context.Context, user1, user2 int64) ([]models.Message, error) {
	return s.storage.List(ctx, user1, user2)
}
