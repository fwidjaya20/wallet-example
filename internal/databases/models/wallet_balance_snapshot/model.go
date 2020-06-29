package wallet_balance_snapshot

type WalletBalanceSnapshot struct {
	Id          string
	WalletId    string
	Balance     float64
	LastEventId string
}
