package server

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"strconv"

	handler_auth "github.com/dijer/otus-highload/backend/internal/handlers/auth"
	handler_profile "github.com/dijer/otus-highload/backend/internal/handlers/profile"
	handler_user "github.com/dijer/otus-highload/backend/internal/handlers/user"
	middleware_auth "github.com/dijer/otus-highload/backend/internal/middlewares/auth"
	service_user "github.com/dijer/otus-highload/backend/internal/services/user"
	storage_user "github.com/dijer/otus-highload/backend/internal/storage/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/dijer/otus-highload/backend/internal/config"
)

type Server struct {
	cfg     config.ServerConf
	db      *sql.DB
	authCfg config.AuthConf
}

func New(cfg config.ServerConf, db *sql.DB, authCfg config.AuthConf) *Server {
	return &Server{
		cfg:     cfg,
		db:      db,
		authCfg: authCfg,
	}
}

func (s *Server) Start(ctx context.Context) error {
	r := mux.NewRouter()

	userStorage := storage_user.New(s.db)
	userService := service_user.New(userStorage)

	authHandler := handler_auth.New(userService, s.authCfg)
	r.HandleFunc("/login", authHandler.LoginHandler)
	r.HandleFunc("/user/register", authHandler.RegisterHandler)
	authMiddleware := middleware_auth.New(s.authCfg)
	r.Handle("/user/logout", authMiddleware.Handler(http.HandlerFunc(authHandler.LogoutHandler)))
	r.Handle("/user/check", authMiddleware.Handler(http.HandlerFunc(authHandler.CheckAuthHandler)))

	userHandler := handler_user.New(userService)
	r.HandleFunc("/user/get/{id}", userHandler.Handler)

	profileHandler := handler_profile.New(userService)
	r.Handle("/user/profile", authMiddleware.Handler(http.HandlerFunc(profileHandler.Handler)))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	println("Server runs on port: " + strconv.Itoa(s.cfg.Port))
	return http.ListenAndServe(net.JoinHostPort(s.cfg.Host, strconv.Itoa(s.cfg.Port)), corsHandler.Handler(r))
}
