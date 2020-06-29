package wallet_balance_event

type Model struct {
	Id            string
	WalletId      string
	TransactionId string
	Amount        float64
	BalanceType   string
	Notes         string
}
