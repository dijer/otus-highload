package middleware_auth

import (
	"context"
	"net/http"

	cache_auth "github.com/dijer/otus-highload/backend/internal/cache/auth"
	"github.com/dijer/otus-highload/backend/internal/config"
	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	cfg   config.AuthConf
	cache *cache_auth.AuthCache
}

func New(cfg config.AuthConf, cache *cache_auth.AuthCache) *AuthMiddleware {
	return &AuthMiddleware{
		cfg:   cfg,
		cache: cache,
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

		uuid, ok := claims["uuid"].(string)
		if !ok {
			utils_server.JsonError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		userID, err := m.cache.GetSession(r.Context(), uuid)
		if err != nil {
			utils_server.JsonError(w, http.StatusInternalServerError, "auth cache get session error")
			return
		}
		if userID == 0 {
			utils_server.JsonError(w, http.StatusUnauthorized, "session expired")
			return
		}

		ctx := context.WithValue(r.Context(), httpctx.ContextUserIDKey, int(userID))
		ctx = context.WithValue(ctx, httpctx.ContextSessionUUIDKey, uuid)

		next.ServeHTTP(w, r.WithContext((ctx)))
	})
}
