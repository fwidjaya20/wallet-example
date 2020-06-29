package chi

import (
	"context"
	"net/http"
	"time"
)

const (
	KEY_REQUEST_TIME = "key_request_time"
)

func RequestTime(next http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, KEY_REQUEST_TIME, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}