package server_websocket

import (
	"context"
	"net"
	"net/http"
	"strconv"

	cache_feed "github.com/dijer/otus-highload/backend/internal/cache/feed"
	"github.com/dijer/otus-highload/backend/internal/config"
	handler_posts_subscribe "github.com/dijer/otus-highload/backend/internal/handlers/posts-subscribe"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/logger"
	middleware_auth "github.com/dijer/otus-highload/backend/internal/middlewares/auth"
	service_posts "github.com/dijer/otus-highload/backend/internal/services/posts"
	storage_posts "github.com/dijer/otus-highload/backend/internal/storage/posts"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
)

type WsServer struct {
	cfg      config.ServerConf
	dbRouter infra_database.DBRouter
	authCfg  config.AuthConf
	log      logger.Logger
	redis    *redis.Client
	rabbitmq *amqp.Channel
}

func New(
	cfg config.ServerConf,
	dbRouter infra_database.DBRouter,
	authCfg config.AuthConf,
	log logger.Logger,
	redis *redis.Client,
	rabbitmq *amqp.Channel,
) *WsServer {
	return &WsServer{
		cfg:      cfg,
		dbRouter: dbRouter,
		authCfg:  authCfg,
		log:      log,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}

func (s *WsServer) Start(ctx context.Context) error {
	r := mux.NewRouter()

	authMiddleware := middleware_auth.New(s.authCfg)

	postsCache := cache_feed.New(s.redis)
	postsStorage := storage_posts.New(s.dbRouter, postsCache, s.log)
	postsService := service_posts.New(postsStorage, s.rabbitmq)

	postsSuscribeHandler := handler_posts_subscribe.New(postsService, s.rabbitmq)
	r.Handle("/post/feed/posted", authMiddleware.Handler(http.HandlerFunc(postsSuscribeHandler.SubscribeFeed))).Methods(http.MethodGet)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	s.log.Info("Websocket Server runs on port: " + strconv.Itoa(s.cfg.WSPort))
	return http.ListenAndServe(net.JoinHostPort(s.cfg.Host, strconv.Itoa(s.cfg.WSPort)), corsHandler.Handler(r))
}
