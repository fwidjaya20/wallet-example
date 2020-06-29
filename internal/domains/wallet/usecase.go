package wallet

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_event"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
)

type UseCase interface {
	Deposit(ctx context.Context, payload models.TopUpRequest) error
	GetBalance(ctx context.Context, payload models.GetBalanceRequest) (*models.Balance, error)
	GetTransaction(ctx context.Context, payload models.GetTransactionEvent) ([]*wallet_balance_event.Model, error)
}