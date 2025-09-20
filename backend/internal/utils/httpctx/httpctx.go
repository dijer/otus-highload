package httpctx

import "net/http"

type ctxKey string

const UserIDKey ctxKey = "userId"

func GetUserID(r *http.Request) int64 {
	val := r.Context().Value(UserIDKey)
	if v, ok := val.(int64); ok {
		return v
	}

	return 0
}
