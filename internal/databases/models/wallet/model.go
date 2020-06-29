package wallet

type Model struct {
	Id string
	Account WalletAccount
}

type WalletAccount struct {
	Holder string
	Number string
}