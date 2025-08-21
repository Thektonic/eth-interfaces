package transaction

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxOptsBuilderFunc func() (*bind.TransactOpts, error)

func Transact[T any](interaction Interaction, s Session, calldata []byte, unpack func([]byte) (T, error)) (*types.Transaction, error) {
	if interaction.Safe() {
		_, err := Call(s, calldata, unpack)
		if err != nil {
			fmt.Println(err.Error())
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
