package repositories

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_event"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_snapshot"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
	"github.com/google/uuid"
)

type Interface interface {
	StoreEvent(ctx context.Context, model *wallet_balance_event.Model) error
	GetBalance(ctx context.Context, walletId uuid.UUID) (*models.Balance, error)
	GetEvents(ctx context.Context, transactionId uuid.UUID) ([]*wallet_balance_snapshot.Model, error)
}