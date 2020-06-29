package wallet_balance_event

type Model struct {
	Id            string  `db:"id"`
	WalletId      string  `db:"wallet_id"`
	TransactionId string  `db:"transaction_id"`
	Amount        float64 `db:"amount"`
	BalanceType   string  `db:"balance_type"`
	Notes         string  `db:"notes"`
}
