package models

type Balance struct {
	Balance float64 `json:"balance"`
}

type TransactionEvent struct {
	TransactionId string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	BalanceType   string  `json:"balance_type"`
	Notes         string  `json:"notes"`
}