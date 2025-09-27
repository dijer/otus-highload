package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/dijer/otus-highload/backend/internal/config"
	databus_feed "github.com/dijer/otus-highload/backend/internal/databus/feed"
	infra_citus "github.com/dijer/otus-highload/backend/internal/infra/citus"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/infra/rabbitmq"
	infra_redis "github.com/dijer/otus-highload/backend/internal/infra/redis"
	"github.com/dijer/otus-highload/backend/internal/logger"
	server_websocket "github.com/dijer/otus-highload/backend/internal/server/websocket"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger := logger.New(log)

	cfg, err := config.New(configFile)
	if err != nil {
		log.Error(err.Error())
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	dbRouter, err := infra_database.New(ctx, cfg.Database, cfg.Replicas)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer dbRouter.Close()

	if cfg.Citus.Coordinator {
		if err := infra_citus.InitCitus(ctx, dbRouter, cfg.Citus.Nodes, cfg.Database); err != nil {
			log.Fatal(err.Error())
		}
	}

	redis, err := infra_redis.InitRedis(ctx, cfg.Redis)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer redis.Close()

	rConn, rCh, err := rabbitmq.New(cfg.RabbitMQ)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer rConn.Close()
	defer rCh.Close()

	if err := databus_feed.ConsumePostCreated(ctx, rCh, dbRouter); err != nil {
		log.Error(err.Error())
		return
	}

	server := server_websocket.New(cfg.Server, *dbRouter, cfg.Auth, logger, redis, rCh)
	errCh := make(chan error, 1)

	go func() {
		errCh <- server.Start(ctx)
	}()

	select {
	case <-ctx.Done():
		log.Info("Shutdown signal")
	case err := <-errCh:
		if err != nil {
			log.Error("Server error: " + err.Error())
		}
	}

	log.Info("Server stopped")
}
