package models

type TopUpRequest struct {
	WalletId string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}

type GetBalanceRequest struct {
	WalletId string `http_url:"wallet_id"`
}

type GetTransactionEvent struct {
	TransactionId string `http_url:"transaction_id"`
}