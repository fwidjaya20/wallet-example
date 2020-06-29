package endpoints

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
	"github.com/fwidjaya20/wallet-example/internal/globals"
	"github.com/fwidjaya20/wallet-example/lib/database"
	libHttp "github.com/fwidjaya20/wallet-example/lib/transport/http"
	"github.com/go-kit/kit/endpoint"
)

func TopUp(service wallet.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		payload := request.(*models.TopUpRequest)

		err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
			err = service.Deposit(ctx, *payload)
			return err
		})

		return libHttp.Response(ctx, response, nil), err
	}
}