package infra_database

import (
	"context"
	"database/sql"
	"fmt"
	"sync/atomic"

	"github.com/dijer/otus-highload/backend/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type DBRouter struct {
	Master   *sql.DB
	Replicas []*sql.DB
	counter  int32
}

func New(ctx context.Context, config config.DatabaseConf, replicasConfig []config.DatabaseConf) (*DBRouter, error) {
	masterDb, err := sql.Open("postgres", buildDsn(config))
	if err != nil {
		return nil, err
	}

	if err := goose.Up(masterDb, "./migrations"); err != nil {
		return nil, err
	}

	err = masterDb.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	var replicas []*sql.DB
	for _, replicaConfig := range replicasConfig {
		replicaDb, err := sql.Open("postgres", buildDsn(replicaConfig))
		if err != nil {
			continue
		}
		if err := replicaDb.PingContext(ctx); err != nil {
			continue
		}
		replicas = append(replicas, replicaDb)
	}

	return &DBRouter{
		Master:   masterDb,
		Replicas: replicas,
	}, nil
}

func (r *DBRouter) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return r.Master.ExecContext(ctx, query, args...)
}

func (r *DBRouter) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if len(r.Replicas) == 0 {
		return r.Master.QueryContext(ctx, query, args...)
	}

	idx := int(atomic.AddInt32(&r.counter, 1)) % len(r.Replicas)
	replica := r.Replicas[idx]

	return replica.QueryContext(ctx, query, args...)
}

func (r *DBRouter) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	replica := r.pickReplica()
	return replica.QueryRowContext(ctx, query, args...)
}

func (r *DBRouter) Close() error {
	if err := r.Master.Close(); err != nil {
		return err
	}
	for _, replica := range r.Replicas {
		if err := replica.Close(); err != nil {
			return err
		}
	}
	return nil
}

func buildDsn(c config.DatabaseConf) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DBName,
	)
}

func (r *DBRouter) pickReplica() *sql.DB {
	if len(r.Replicas) == 0 {
		return r.Master
	}
	idx := int(atomic.AddInt32(&r.counter, 1)) % len(r.Replicas)
	return r.Replicas[idx]
}
