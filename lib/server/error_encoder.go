package server

import (
	"context"
	"encoding/json"
	libError "github.com/fwidjaya20/wallet-example/lib/error"
	"net/http"
)

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := http.StatusInternalServerError
	message := "Something Went Wrong"

	if sc, ok := err.(*libError.Error); ok {
		code = sc.StatusCode
		message = sc.Message
	}

	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error":  message,
		"message": err.Error(),
	})
}
