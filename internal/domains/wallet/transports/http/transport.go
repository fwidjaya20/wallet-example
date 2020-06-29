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

func Withdraw(service wallet.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.Withdraw(service), server.HTTPOption{
			DecodeModel: &models.Withdraw{},
		}, opts).ServeHTTP(w, r)
	}
}

func GetBalance(service wallet.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.GetBalance(service), server.HTTPOption{
			DecodeModel: &models.GetBalanceRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func GetTransaction(service wallet.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.GetTransaction(service), server.HTTPOption{
			DecodeModel: &models.GetTransactionEvent{},
		}, opts).ServeHTTP(w, r)
	}
}
