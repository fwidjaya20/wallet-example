package repositories

import (
	"context"
	"database/sql"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_event"
	"github.com/fwidjaya20/wallet-example/internal/databases/models/wallet_balance_snapshot"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/models"
	"github.com/fwidjaya20/wallet-example/internal/globals"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type postgres struct {}

func (p *postgres) StoreEvent(ctx context.Context, model *wallet_balance_event.Model) error {
	var err error
	query, args := p.buildStoreEventQuery(model)
	_, err = globals.GetQuery(ctx).NamedExecContext(ctx, query, args)

	if nil != err {
		return err
	}

	err = p.createSnapshot(ctx, model)

	return err
}

func (p *postgres) GetBalance(ctx context.Context, walletId string) (*models.Balance, error) {
	snapshot, err := p.getLastSnapshot(ctx, walletId)

	if nil != err {
		return nil, err
	}

	return &models.Balance{
		Balance: snapshot.Balance,
	}, nil
}

func (p *postgres) GetEvents(ctx context.Context, walletId string) ([]*wallet_balance_event.Model, error) {
	var result []*wallet_balance_event.Model
	var err error
	var rows *sqlx.Rows

	query, args := p.buildGetEventQuery(walletId)
	rows, err = globals.GetQuery(ctx).NamedQueryxContext(ctx, query, args)

	if nil != err {
		return nil, err
	}

	for rows.Next() {
		var model wallet_balance_event.Model
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	_ = rows.Close()

	return result, err
}

func NewWalletRepository() Interface {
	return &postgres{}
}

func (p *postgres) createSnapshot(ctx context.Context, model *wallet_balance_event.Model) error {
	lastSnapshot, err := p.getLastSnapshot(ctx, model.WalletId)
	if nil != err && err != sql.ErrNoRows {
		return err
	}

	snapshotId, err := uuid.NewRandom()
	if nil != err {
		return err
	}

	snapshotModel := wallet_balance_snapshot.Model{
		Id:          snapshotId.String(),
		WalletId:    model.WalletId,
		Balance:     lastSnapshot.Balance + model.Amount,
		LastEventId: model.Id,
	}

	query, args := p.buildCreateSnapshotQuery(snapshotModel)
	_, err = globals.GetQuery(ctx).NamedExecContext(ctx, query, args)

	return err
}

func (p *postgres) getLastSnapshot(ctx context.Context, walletId string) (*wallet_balance_snapshot.Model, error) {
	var result wallet_balance_snapshot.Model
	var err error

	query, arg := p.buildGetLastSnapshotQuery(walletId)
	row, err := globals.GetQuery(ctx).NamedQueryRowxContext(ctx, query, arg)

	if nil != err {
		return nil, err
	}

	err = row.StructScan(&result)

	return &result, err
}

func (p *postgres) buildGetLastSnapshotQuery(walletId string) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`SELECT "id", "wallet_id", "balance", "last_event_id" `)
	query.WriteString(`FROM "wallet_balance_snapshots" `)
	query.WriteString(`WHERE "wallet_id"=:walletId `)
	query.WriteString(`ORDER BY "created_at" DESC `)
	query.WriteString(`LIMIT 1`)

	arg["walletId"] = walletId

	return query.String(), arg
}

func (p *postgres) buildStoreEventQuery(model *wallet_balance_event.Model) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`INSERT INTO "wallet_balance_events" `)
	query.WriteString(`("id", "wallet_id", "transaction_id", "amount", "balance_type", "notes", "created_by") `)
	query.WriteString(`VALUES `)
	query.WriteString(`(:id, :walletId, :transactionId, :amount, :balanceType, :notes, :createdBy)`)

	arg["id"] = model.Id
	arg["walletId"] = model.WalletId
	arg["amount"] = model.Amount
	arg["balanceType"] = model.BalanceType
	arg["notes"] = model.Notes
	arg["transactionId"] = model.TransactionId
	arg["createdBy"] = "SYSTEM"

	return query.String(), arg
}

func (p *postgres) buildCreateSnapshotQuery(model wallet_balance_snapshot.Model) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`INSERT INTO "wallet_balance_snapshots" `)
	query.WriteString(`("id", "wallet_id", "balance", "last_event_id", "created_by") `)
	query.WriteString(`VALUES `)
	query.WriteString(`(:id, :walletId, :balance, :lastEventId, :createdBy)`)

	arg["id"] = model.Id
	arg["walletId"] = model.WalletId
	arg["balance"] = model.Balance
	arg["lastEventId"] = model.LastEventId
	arg["createdBy"] = "SYSTEM"

	return query.String(), arg
}

func (p *postgres) buildGetEventQuery(walletId string) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`SELECT "id", "wallet_id", "transaction_id", "amount", "balance_type", "notes" `)
	query.WriteString(`FROM "wallet_balance_events" `)
	query.WriteString(`WHERE "wallet_id"=:walletId `)
	query.WriteString(`ORDER BY "created_at" DESC `)

	arg["walletId"] = walletId

	return query.String(), arg
}