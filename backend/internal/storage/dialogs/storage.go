package storage_dialogs

import (
	"context"
	"database/sql"
	"errors"

	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/models"
)

type DialogsStorage struct {
	dbRouter infra_database.DBRouter
}

func New(dbRouter infra_database.DBRouter) *DialogsStorage {
	return &DialogsStorage{
		dbRouter: dbRouter,
	}
}

func (s *DialogsStorage) GetChatID(ctx context.Context, user1, user2 int64) (int64, error) {
	if user1 == user2 {
		return 0, errors.New("user1 equals user2")
	}

	a, b := user1, user2
	if a > b {
		a, b = b, a
	}

	var chatID int64
	err := s.dbRouter.QueryRow(ctx, `
		SELECT chat_id FROM dialogs
		WHERE user_a=$1 AND user_b=$2
	`, a, b).Scan(&chatID)

	if err == sql.ErrNoRows {
		err = s.dbRouter.QueryRow(ctx, `
			INSERT INTO dialogs (user_a, user_b)
			VALUES ($1, $2)
			RETURNING chat_id
		`, a, b).Scan(&chatID)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (s *DialogsStorage) Send(ctx context.Context, user1, user2 int64, text string) error {
	chatID, err := s.GetChatID(ctx, user1, user2)
	if err != nil {
		return err
	}

	_, err = s.dbRouter.Exec(ctx, `
		INSERT INTO messages(chat_id, sender_id, recipient_id, body)
		VALUES ($1, $2, $3, $4)
	`, chatID, user1, user2, text)

	return err
}

func (s *DialogsStorage) List(ctx context.Context, user1, user2 int64) ([]models.Message, error) {
	chatID, err := s.GetChatID(ctx, user1, user2)
	if err != nil {
		return nil, err
	}

	rows, err := s.dbRouter.Query(ctx, `
		SELECT sender_id, recipient_id, body
		FROM messages
		WHERE chat_id = $1
		ORDER BY msg_id ASC
	`, chatID)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.From, &m.To, &m.Text); err != nil {
			return nil, err
		}

		messages = append(messages, m)
	}

	return messages, nil
}
