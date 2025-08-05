// Package hex provides common utilities and constants for Ethereum contract interactions.
package hex

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// GetFunctionSelector returns the function selector for a given signature
func GetFunctionSelector(signature Signature) string {
	return signature.GetHex()
}

// GetEncodedFunction encodes a function call with the given parameters
func GetEncodedFunction(abiString, signature string, params ...interface{}) ([]byte, error) {
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	return contractABI.Pack(signature, params...)
}

// DeployContract deploys a contract with the given parameters
func DeployContract(auth *bind.TransactOpts,
	client simulated.Client,
	abiString string,
	byteCodeString string,
	params ...interface{},
) (common.Address, *ethTypes.Transaction, *bind.BoundContract, error) {
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var byteCode []byte
	if byteCodeString[0:2] == "0x" {
		byteCode = common.Hex2Bytes(byteCodeString[2:])
	} else {
		byteCode = common.Hex2Bytes(byteCodeString)
	}

	return bind.DeployContract(
		auth,
		contractABI,
		byteCode,
		client,
		params...,
	)
}
