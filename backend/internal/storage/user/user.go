package storage_user

import (
	"database/sql"

	errs "github.com/dijer/otus-highload/backend/internal/errors"
	"github.com/dijer/otus-highload/backend/internal/models"
	"github.com/lib/pq"
)

type UserStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) CreateUser(user models.User, hashedPassword string) error {
	_, err := s.db.Exec(`INSERT INTO users (username, password_hash, first_name, last_name, birthday, gender, interests, city)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.UserName,
		hashedPassword,
		user.FirstName,
		user.LastName,
		user.Birthday,
		user.Gender,
		pq.Array(user.Interests),
		user.City,
	)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			return errs.ErrUserAlreadyExists
		}
	}

	return err
}

func (s *UserStorage) GetHashedPassword(username string) (string, int, error) {
	var hashedPassword string
	var userID int
	err := s.db.QueryRow(`SELECT id, password_hash FROM users WHERE username = $1`, username).Scan(&userID, &hashedPassword)
	if err != nil {
		return "", 0, err
	}

	return hashedPassword, userID, nil
}

func (s *UserStorage) GetUser(userID int) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow(`SELECT username, first_name, last_name, birthday, gender, interests, city FROM users WHERE id = $1`, userID).Scan(
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.Birthday,
		&user.Gender,
		pq.Array(&user.Interests),
		&user.City,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
