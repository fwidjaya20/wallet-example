package wallet

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
)

type UseCase interface {
	Deposit(ctx context.Context, payload models.TopUpRequest) error
	GetBalance(ctx context.Context, payload models.GetBalanceRequest) (*models.Balance, error)
	GetTransaction(ctx context.Context, payload models.GetTransactionEvent) ([]*models.TransactionEvent, error)
	Withdraw(ctx context.Context, payload models.Withdraw) error
}