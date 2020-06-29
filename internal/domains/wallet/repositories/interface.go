package repositories

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_event"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
)

type Interface interface {
	StoreEvent(ctx context.Context, model *wallet_balance_event.Model) error
	GetBalance(ctx context.Context, walletId string) (*models.Balance, error)
	GetEvents(ctx context.Context, walletId string) ([]*wallet_balance_event.Model, error)
}