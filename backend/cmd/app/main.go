package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dijer/otus-highload/backend/internal/config"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/server"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.New(configFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	db, err := infra_database.New(ctx, cfg.Database)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	server := server.New(cfg.Server, db, cfg.Auth)

	go func() {
		err = server.Start(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	<-ctx.Done()

	println("Server stopped")
}
