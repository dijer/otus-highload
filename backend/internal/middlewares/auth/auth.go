package middleware_auth

import (
	"context"
	"net/http"

	"github.com/dijer/otus-highload/backend/internal/config"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	cfg config.AuthConf
}

func New(cfg config.AuthConf) *AuthMiddleware {
	return &AuthMiddleware{
		cfg: cfg,
	}
}

func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(m.cfg.JWTCookieName)
		if err != nil {
			utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(m.cfg.JWTKey), nil
		})

		if err != nil || !token.Valid {
			utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils_server.JsonError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		userID, ok := claims["userId"].(float64)
		if !ok {
			utils_server.JsonError(w, http.StatusUnauthorized, "Invalid subject in token")
			return
		}

		ctx := context.WithValue(r.Context(), "userId", int(userID))

		next.ServeHTTP(w, r.WithContext((ctx)))
	})
}
