package server

import (
	"context"
	"net"
	"net/http"
	"strconv"

	cache_feed "github.com/dijer/otus-highload/backend/internal/cache/feed"
	handler_auth "github.com/dijer/otus-highload/backend/internal/handlers/auth"
	handler_friend "github.com/dijer/otus-highload/backend/internal/handlers/friend"
	handler_posts "github.com/dijer/otus-highload/backend/internal/handlers/posts"
	handler_profile "github.com/dijer/otus-highload/backend/internal/handlers/profile"
	handler_user "github.com/dijer/otus-highload/backend/internal/handlers/user"
	handler_user_search "github.com/dijer/otus-highload/backend/internal/handlers/user-search"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/logger"
	middleware_auth "github.com/dijer/otus-highload/backend/internal/middlewares/auth"
	service_friend "github.com/dijer/otus-highload/backend/internal/services/friend"
	service_posts "github.com/dijer/otus-highload/backend/internal/services/posts"
	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
	storage_friend "github.com/dijer/otus-highload/backend/internal/storage/friend"
	storage_posts "github.com/dijer/otus-highload/backend/internal/storage/posts"
	storage_user "github.com/dijer/otus-highload/backend/internal/storage/user"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"

	"github.com/dijer/otus-highload/backend/internal/config"
)

type Server struct {
	cfg      config.ServerConf
	dbRouter infra_database.DBRouter
	authCfg  config.AuthConf
	log      logger.Logger
	redis    *redis.Client
}

func New(
	cfg config.ServerConf,
	dbRouter infra_database.DBRouter,
	authCfg config.AuthConf,
	log logger.Logger,
	redis *redis.Client,
) *Server {
	return &Server{
		cfg:      cfg,
		dbRouter: dbRouter,
		authCfg:  authCfg,
		log:      log,
	}
}

func (s *Server) Start(ctx context.Context) error {
	r := mux.NewRouter()

	userStorage := storage_user.New(s.dbRouter)
	userService := service_user.New(userStorage)

	authHandler := handler_auth.New(userService, s.authCfg)
	r.HandleFunc("/login", authHandler.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/register", authHandler.RegisterHandler).Methods(http.MethodPost)
	authMiddleware := middleware_auth.New(s.authCfg)
	r.Handle("/user/logout", authMiddleware.Handler(http.HandlerFunc(authHandler.LogoutHandler))).Methods(http.MethodGet)
	r.Handle("/user/check", authMiddleware.Handler(http.HandlerFunc(authHandler.CheckAuthHandler))).Methods(http.MethodGet)

	userHandler := handler_user.New(userService)
	r.HandleFunc("/user/get/{id}", userHandler.Handler).Methods(http.MethodGet)

	profileHandler := handler_profile.New(userService)
	r.Handle("/user/profile", authMiddleware.Handler(http.HandlerFunc(profileHandler.Handler))).Methods(http.MethodGet)

	userSearchHandler := handler_user_search.New(userService)
	r.HandleFunc("/user/search", userSearchHandler.Handler).Methods(http.MethodGet)

	friendStorage := storage_friend.New(s.dbRouter)
	friendService := service_friend.New(friendStorage)
	friendHandler := handler_friend.New(friendService)
	r.Handle("/friend/set/{userId}", authMiddleware.Handler(http.HandlerFunc(friendHandler.AddFriend))).Methods(http.MethodPut)
	r.Handle("/friend/delete/{userId}", authMiddleware.Handler(http.HandlerFunc(friendHandler.RemoveFriend))).Methods(http.MethodPut)

	postsCache := cache_feed.New(s.redis)
	postsStorage := storage_posts.New(s.dbRouter, postsCache, s.log)
	postsService := service_posts.New(postsStorage)
	postsHandler := handler_posts.New(postsService)
	r.Handle("/post/create", authMiddleware.Handler(http.HandlerFunc(postsHandler.CreatePost))).Methods(http.MethodPost)
	r.Handle("/post/feed", authMiddleware.Handler(http.HandlerFunc(postsHandler.GetFeed))).Methods(http.MethodGet)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	s.log.Info("Server runs on port: " + strconv.Itoa(s.cfg.Port))
	return http.ListenAndServe(net.JoinHostPort(s.cfg.Host, strconv.Itoa(s.cfg.Port)), corsHandler.Handler(r))
}
