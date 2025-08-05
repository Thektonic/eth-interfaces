// Package contractextension provides interfaces and utilities for extending contract functionality.
package contractextension

import (
	"context"

	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/ethereum/go-ethereum/common"
)

// ContractImplementation defines the interface for contract implementation verification
type ContractImplementation interface {
	GetAddress() common.Address
	VerifyTransaction(ctx context.Context, to common.Address, data []byte, value int64) error
}

// SimulateCall simulates a contract call to verify implementation
func SimulateCall(ctx context.Context,
	abiPath,
	selector string,
	contract ContractImplementation,
	params ...interface{},
) error {
	encodedFunction, err := hex.GetEncodedFunction(abiPath, selector, params...)
	if err != nil {
		return err
	}
	return contract.VerifyTransaction(ctx, contract.GetAddress(), encodedFunction, 0)
}
