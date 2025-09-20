package storage_user

import (
	"context"

	errs "github.com/dijer/otus-highload/backend/internal/errors"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/models"
	"github.com/lib/pq"
)

type UserStorage struct {
	dbRouter infra_database.DBRouter
}

func New(dbRouter infra_database.DBRouter) *UserStorage {
	return &UserStorage{
		dbRouter: dbRouter,
	}
}

func (s *UserStorage) CreateUser(ctx context.Context, user models.User, hashedPassword string) error {
	_, err := s.dbRouter.Exec(ctx, `INSERT INTO users (username, password_hash, first_name, last_name, birthday, gender, interests, city)
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

func (s *UserStorage) GetHashedPassword(ctx context.Context, username string) (string, int64, error) {
	var hashedPassword string
	var userID int64
	err := s.dbRouter.QueryRow(ctx, `SELECT id, password_hash FROM users WHERE username = $1`, username).Scan(&userID, &hashedPassword)
	if err != nil {
		return "", 0, err
	}

	return hashedPassword, userID, nil
}

func (s *UserStorage) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	err := s.dbRouter.QueryRow(ctx, `SELECT username, first_name, last_name, birthday, gender, interests, city FROM users WHERE id = $1`, userID).Scan(
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

func (s *UserStorage) GetUsers(ctx context.Context, firstname, lastname string) ([]models.User, error) {
	users := make([]models.User, 0)

	firstnamePattern := firstname + "%"
	lastnamePattern := lastname + "%"

	rows, err := s.dbRouter.Query(ctx, `
		SELECT username, first_name, last_name, birthday, gender, interests, city
		FROM users
		WHERE LOWER(first_name) LIKE LOWER($1) AND LOWER(last_name) LIKE LOWER($2)
		ORDER BY id`,
		firstnamePattern, lastnamePattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
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

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
