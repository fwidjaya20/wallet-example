package wallet_balance_event

type WalletBalanceEvent struct {
	Id          string
	WalletId    string
	Amount      float64
	BalanceType string
	Notes       string
}
