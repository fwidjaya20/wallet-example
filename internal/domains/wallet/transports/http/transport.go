package http

import (
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/endpoints"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
	"github.com/fwidjaya20/wallet-example/lib/server"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func TopUp(service wallet.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.TopUp(service), server.HTTPOption{
			DecodeModel: &models.TopUpRequest{},
		}, opts).ServeHTTP(w, r)
	}
}
