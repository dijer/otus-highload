package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/dijer/otus-highload/backend/internal/config"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/logger"
	"github.com/dijer/otus-highload/backend/internal/server"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.toml", "Path to configuration file")
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

	server := server.New(cfg.Server, *dbRouter, cfg.Auth, logger)
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
