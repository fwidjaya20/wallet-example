package wallet

import (
	"context"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_event"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/repositories"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/values/balance"
	libError "github.com/fwidjaya20/wallet-example/lib/error"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"net/http"
)

type service struct {
	actor string
	logger log.Logger
	repository repositories.Interface
}

func (s *service) Deposit(ctx context.Context, payload models.TopUpRequest) error {
	logger := log.With(s.logger, "METHOD", "Deposit()")

	var err error

	balanceId, err := uuid.NewRandom()
	if nil != err {
		_ = level.Error(logger).Log("create_uuid_failed", err)
		return libError.NewError(err, http.StatusInternalServerError, "create_uuid_failed")
	}

	transId, err := uuid.NewRandom()
	if nil != err {
		_ = level.Error(logger).Log("create_uuid_failed", err)
		return libError.NewError(err, http.StatusInternalServerError, "create_uuid_failed")
	}

	err = s.repository.StoreEvent(ctx, &wallet_balance_event.Model{
		Id:            balanceId.String(),
		WalletId:      payload.WalletId,
		TransactionId: transId.String(),
		Amount:        payload.Amount,
		BalanceType:   balance.DEPOSIT,
		Notes:         `{"message": "add"}`,
	})

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return libError.NewError(err, http.StatusInternalServerError, "create_deposit_error")
	}

	return nil
}

func (s *service) GetBalance(ctx context.Context, payload models.GetBalanceRequest) (*models.Balance, error) {
	logger := log.With(s.logger, "METHOD", "GetBalance()")

	var result *models.Balance
	var err error

	result, err = s.repository.GetBalance(ctx, payload.WalletId)

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return nil, libError.NewError(err, http.StatusInternalServerError, "get_wallet_balance_error")
	}

	return result, nil
}

func (s *service) GetTransaction(ctx context.Context, payload models.GetTransactionEvent) ([]*wallet_balance_event.Model, error) {
	logger := log.With(s.logger, "METHOD", "GetEvent()")

	var result []*wallet_balance_event.Model
	var err error

	result, err = s.repository.GetEvents(ctx, payload.WalletId)

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return nil, libError.NewError(err, http.StatusInternalServerError, "get_wallet_event_error")
	}

	return result, nil
}

func NewWalletService(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor:      "WALLET",
		logger:     nil,
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}