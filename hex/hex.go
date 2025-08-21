// Package hex provides common utilities and constants for Ethereum contract interactions.
package hex

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

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

// GetImplementationAddress retrieves the implementation address from an EIP-1967 proxy contract
func GetImplementationAddress(
	ctx context.Context,
	client simulated.Client,
	proxyAddr common.Address,
) (common.Address, error) {
	// EIP-1967 implementation slot
	slot := common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	result, err := client.StorageAt(ctx, proxyAddr, slot, nil)
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(result), nil
}

// CheckDiamondFunction checks if a function selector is supported by a Diamond proxy contract
func CheckDiamondFunction(
	ctx context.Context,
	client simulated.Client,
	diamondAddr common.Address,
	funcSelector []byte,
) (bool, error) {
	var selector [4]byte
	copy(selector[:], funcSelector)

	data, err := GetEncodedFunction(diamondLoupeABI, "facetAddress", selector)
	if err != nil {
		return false, fmt.Errorf("failed to encode facetAddress call: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   &diamondAddr,
		Data: data,
	}

	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return false, err
	}

	return common.BytesToAddress(result) != (common.Address{}), nil
}

// DecodeErrorData decodes hex-encoded error data from various input formats.
func DecodeErrorData(data interface{}) []byte {
	switch v := data.(type) {
	case string:
		// Remove "0x" prefix if present
		if len(v) >= 2 && v[:2] == "0x" {
			return common.Hex2Bytes(v[2:])
		}
		return common.Hex2Bytes(v)
	case []byte:
		return v
	}
	return nil
}
