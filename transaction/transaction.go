package transaction

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

type TxOptsBuilderFunc func() (*bind.TransactOpts, error)
