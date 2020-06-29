package models

type TopUpRequest struct {
	WalletId string  `http_url:"wallet_id"`
	Amount   float64 `json:"amount"`
}

type GetBalanceRequest struct {
	WalletId string `http_url:"wallet_id"`
}

type GetTransactionEvent struct {
	WalletId      string `http_url:"wallet_id"`
	TransactionId string `http_url:"transaction_id"`
}