package http

import (
	"github.com/fwidjaya20/wallet-example/cmd/container"
	walletHttpTransport "github.com/fwidjaya20/wallet-example/internal/domains/wallet/transports/http"
	libServer "github.com/fwidjaya20/wallet-example/lib/server"
	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeHandler(
	router *chi.Mux,
	container container.Container) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(libServer.ErrorEncoder),
	}

	generateWalletRoute(router, container, opts)

	return router
}

func generateWalletRoute(router chi.Router, container container.Container, opts []kitHttp.ServerOption) {
	router.Group(func(r chi.Router) {
		r.Post("/wallet/{wallet_id}/top-up", walletHttpTransport.TopUp(container.WalletService, opts))
		r.Get("/wallet/{wallet_id}/balance", walletHttpTransport.GetBalance(container.WalletService, opts))
		r.Get("/wallet/{wallet_id}/transactions", walletHttpTransport.GetTransaction(container.WalletService, opts))
	})
}