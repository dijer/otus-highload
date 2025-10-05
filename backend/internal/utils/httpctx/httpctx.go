package httpctx

import "net/http"

type contextKey string

const (
	ContextUserIDKey      contextKey = "userId"
	ContextSessionUUIDKey contextKey = "uuid"
)

func GetUserID(r *http.Request) int64 {
	val := r.Context().Value(ContextUserIDKey)
	if v, ok := val.(int64); ok {
		return v
	}

	return 0
}

func GetUUID(r *http.Request) string {
	val := r.Context().Value(ContextSessionUUIDKey)
	if v, ok := val.(string); ok {
		return v
	}

	return ""
}
