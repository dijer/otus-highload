package httpctx

import "net/http"

type ctxKey string

const UserIDKey ctxKey = "userId"

func GetUserID(r *http.Request) int {
	val := r.Context().Value(UserIDKey)
	if v, ok := val.(int); ok {
		return v
	}

	return 0
}
