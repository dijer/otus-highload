package infra_database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dijer/otus-highload/backend/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func New(ctx context.Context, config config.DatabaseConf) (*sql.DB, error) {
	println(config.User, config.Password, config.Host, config.Port, config.DBName)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
