// Package testingtools provides common utilities and constants for Ethereum contract interactions.
package testingtools

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// SetupBlockchain sets up a test blockchain environment with a deployed contract
func SetupBlockchain(
	t *testing.T,
	contractABI string,
	byteCode string,
	params ...interface{},
) (
	*simulated.Backend,
	*bind.TransactOpts,
	*common.Address,
	*ecdsa.PrivateKey,
	error,
) {
	privKey, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(hex.TestChainID))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	testUserAddress := crypto.PubkeyToAddress(privKey.PublicKey)
	alloc := ethTypes.GenesisAlloc{
		testUserAddress: ethTypes.Account{Balance: hex.MaxUint256},
	}
	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(hex.TestGasLimit))

	contractAddr, tx, _, err := hex.DeployContract(
		auth,
		backend.Client(),
		contractABI,
		byteCode,
		params...,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	backend.Commit()

	receipt, err := backend.Client().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil || receipt.Status != 1 {
		t.Fatal("contract deployment failed")
	}

	return backend, auth, &contractAddr, privKey, nil
}

// DeployEmptyContract deploys an empty contract for testing purposes
func DeployEmptyContract(auth *bind.TransactOpts, backend *simulated.Backend) (*common.Address, error) {
	contractAddr, tx, _, err := hex.DeployContract(
		auth,
		backend.Client(),
		"[]",
		"0x6080604052348015600e575f5ffd5b50603e80601a5f395ff3fe60806040525f5ffdfea264697066735822122"+
			"0c29a733ea58e61bf35a384ca562ce9f1aa87686a9e1f404e3efbe9b0e388609e64736f6c634300081c0033",
	)
	if err != nil {
		return nil, err
	}
	backend.Commit()

	receipt, err := backend.Client().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil || receipt.Status != 1 {
		return nil, fmt.Errorf("empty contract deployment failed: %w", err)
	}
	return &contractAddr, nil
}
