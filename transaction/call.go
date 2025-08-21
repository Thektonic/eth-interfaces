// Package transaction provides utilities for blockchain transaction operations.
package transaction

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

// Session defines the methods required for a session in a transaction interaction.
type Session interface {
	CallOpts() *bind.CallOpts
	Instance() *bind.BoundContract
}

// Interaction defines the methods required for a transaction interaction.
type Interaction interface {
	BaseTxSetup() (*bind.TransactOpts, error)
	Safe() bool
}

// Call performs a call to the contract using the provided session and calldata, returning the unpacked result.
func Call[T any](s Session, calldata []byte, unpack func([]byte) (T, error)) (T, error) {
	return bind2.Call(s.Instance(), s.CallOpts(), calldata, unpack)
}
