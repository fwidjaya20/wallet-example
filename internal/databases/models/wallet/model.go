package wallet

type Wallet struct {
	Id string
	Account WalletAccount
}

type WalletAccount struct {
	Holder string
	Number string
}