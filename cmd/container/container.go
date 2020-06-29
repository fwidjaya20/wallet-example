package container

import (
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet"
	"github.com/fwidjaya20/wallet-example/internal/domains/wallet/repositories"
	"github.com/go-kit/kit/log"
)

type Container struct {
	WalletService wallet.UseCase
}

func New(
	logger log.Logger,
) Container {
	return Container{
		WalletService: wallet.NewWalletService(logger, repositories.NewWalletRepository()),
	}
}
