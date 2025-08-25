package transaction

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxOptsMiddlewareFunc defines a function type that returns a transaction options builder.
type TxOptsMiddlewareFunc func(*bind.TransactOpts) (*bind.TransactOpts, error)

// Transact is an abstraction for the bind.Transact function, allowing for a more generic transaction interface.
func Transact[T any](
	interaction Interaction,
	s Session,
	calldata []byte,
	unpack func([]byte,
	) (T, error)) (*types.Transaction, error) {
	if interaction.Safe() {
		_, err := Call(s, calldata, unpack)
		if err != nil {
			return nil, err
		}
	}

	txOpts, err := interaction.BaseTxSetup()
	if err != nil {
		return nil, err
	}
	tx, err := bind2.Transact(s.Instance(), txOpts, calldata)
	if err != nil {
		return nil, err
	}
	return tx, err
}
