package wallet_balance_snapshot

type Model struct {
	Id          string  `db:"id"`
	WalletId    string  `db:"wallet_id"`
	Balance     float64 `db:"balance"`
	LastEventId string  `db:"last_event_id"`
}
