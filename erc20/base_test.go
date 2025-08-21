package erc20_test

// Package nft_test contains tests for NFT interactions defined in base.go.

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/erc20"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/testingtools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// Test_DeploySuccessfully tests if the blockchain setup and contract deployment succeed without errors.
func Test_DeploySuccessfully(t *testing.T) {
	backend, _, _, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	_ = privKey
	assert.Nil(t, err, "failed to create interactions interface, error: %w", err)
	if err := backend.Close(); err != nil {
		t.Logf("failed to close backend: %v", err)
	}
}

// Test_Instantiation verifies that the ERC20 interactions interface is correctly instantiated
// using various contracts, including a valid ERC20 contract, an empty contract, and an NFT contract.
func Test_Instantiation(t *testing.T) {
	backend, auth, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	emptyContract, err := testingtools.DeployEmptyContract(auth, backend)
	if err != nil {
		t.Fatalf("failed to deploy empty contract: %s", err)
	}

	erc721CompleteContract, tx, _, err := hex.DeployContract(
		auth,
		backend.Client(),
		inferences.Ierc721MetaData.ABI,
		inferences.Ierc721MetaData.Bin,
		"MyNFT", // Arg 1: name
		"MNFT",  // Arg 2: symbol
	)
	if err != nil {
		t.Fatalf("failed to deploy ERC721A contract: %s", err)
	}

	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil || receipt.Status != 1 {
		t.Fatalf("failed to deploy ERC20 contract: %s", err)
	}

	testCases := []struct {
		Name           string
		ContractAddr   common.Address
		ExpectedResult string
		ExpectError    bool
		ExpectedError  string
	}{
		{
			Name:           "OK - Successfully instantiated",
			ExpectedResult: "MyNFT",
			ContractAddr:   *contractAddress,
		},
		{
			Name:          "KO - Empty contract doesn't implement interface",
			ExpectError:   true,
			ContractAddr:  *emptyContract,
			ExpectedError: "interface setup error function CheckSignatures, error :",
		},
		{
			Name:          "KO - ERC721A doesn't implement the interface",
			ExpectError:   true,
			ContractAddr:  erc721CompleteContract,
			ExpectedError: "interface setup error function CheckSignatures, error :",
		},
	}

	baseInteractions := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := erc20.NewIERC20Interactions(
				baseInteractions,
				tt.ContractAddr,
				[]erc20.BaseERC20Signature{
					erc20.Name,
					erc20.Symbol,
					erc20.Decimals,
				},
			)
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Contains(t, err.Error(), tt.ExpectedError)
			} else {
				assert.NoError(t, err, "failed to create interactions interface, error: %w", err)
			}
		})
	}
}

// testERC20StringMethod is a helper function to test ERC20 string methods (Name, Symbol) to avoid code duplication
func testERC20StringMethod(
	t *testing.T,
	signature erc20.BaseERC20Signature,
	expectedResult string,
	methodCall func(*erc20.Interactions) (string, error),
) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	testCases := []struct {
		Name           string
		ContractAddr   common.Address
		ExpectedResult string
		ExpectError    bool
		ExpectedError  string
	}{
		{
			Name:           "OK - Successfully get token metadata",
			ExpectedResult: expectedResult,
			ContractAddr:   *contractAddress,
		},
	}

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			session, err := erc20.NewIERC20Interactions(base, tt.ContractAddr, []erc20.BaseERC20Signature{signature})
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Equal(t, tt.ExpectedError, err.Error())
			} else {
				assert.Nil(t, err, "failed to create interactions interface, error: %w", err)
				result, err := methodCall(session)
				assert.Nil(t, err)
				assert.Equal(t, tt.ExpectedResult, result)
			}
		})
	}
}

// Test_Name verifies that the ERC20 contract correctly returns its name.
func Test_Name(t *testing.T) {
	testERC20StringMethod(t, erc20.Name, "TESTToken", func(session *erc20.Interactions) (string, error) {
		return session.Name()
	})
}

// Test_Symbol verifies that the ERC20 contract correctly returns its symbol.
func Test_Symbol(t *testing.T) {
	testERC20StringMethod(t, erc20.Symbol, "TT", func(session *erc20.Interactions) (string, error) {
		return session.Symbol()
	})
}

// Test_TotalSupply verifies that the total supply of NFTs is correctly reported by the contract.
func Test_TotalSupply(t *testing.T) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	testCases := []struct {
		Name           string
		ContractAddr   common.Address
		ExpectedResult *big.Int
		ExpectError    bool
		ExpectedError  string
	}{
		{
			Name:           "OK - Successfully get NFT total supply",
			ExpectedResult: big.NewInt(100_000_000),
			ContractAddr:   *contractAddress,
		},
	}

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			session, err := erc20.NewIERC20Interactions(base, tt.ContractAddr, []erc20.BaseERC20Signature{erc20.TotalSupply})
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Equal(t, tt.ExpectedError, err.Error())
			} else {
				assert.Nil(t, err, "failed to create interactions interface, error: %w", err)
				supply, err := session.TotalSupply()
				assert.Nil(t, err)
				assert.Equal(t, 0, supply.Cmp(big.NewInt(0).Mul(tt.ExpectedResult, big.NewInt(1e18))))
			}
		})
	}
}

// Test_Decimals verifies that the total supply of NFTs is correctly reported by the contract.
func Test_Decimals(t *testing.T) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	testCases := []struct {
		Name           string
		ContractAddr   common.Address
		ExpectedResult uint8
		ExpectError    bool
		ExpectedError  string
	}{
		{
			Name:           "OK - Successfully get NFT decimals",
			ExpectedResult: 18,
			ContractAddr:   *contractAddress,
		},
	}

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			session, err := erc20.NewIERC20Interactions(base, tt.ContractAddr, []erc20.BaseERC20Signature{erc20.Decimals})
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Equal(t, tt.ExpectedError, err.Error())
			} else {
				assert.Nil(t, err, "failed to create interactions interface, error: %w", err)
				decimals, err := session.Decimals()
				assert.Nil(t, err)
				assert.Equal(t, tt.ExpectedResult, decimals)
			}
		})
	}
}

// Test_Transfer tests the transfer functionality and ensures that the token transfer behaves as expected.
func Test_Transfer(t *testing.T) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}

	type transferArgs struct {
		pk  *ecdsa.PrivateKey
		To  common.Address
		qty *big.Int
	}

	testCases := []struct {
		Name          string
		ContractAddr  common.Address
		args          transferArgs
		ExpectError   bool
		ExpectedError string
	}{
		{
			Name: "OK - Successfully get transfer NFT",
			args: transferArgs{
				To:  common.HexToAddress("1"),
				qty: big.NewInt(10),
			},
			ContractAddr: *contractAddress,
		},
		{
			Name: "KO - Burn NFT",
			args: transferArgs{
				To:  common.HexToAddress("0"),
				qty: big.NewInt(1),
			},
			ContractAddr:  *contractAddress,
			ExpectError:   true,
			ExpectedError: "call error on erc20.Transfer(): ERC20InvalidReceiver",
		},
		{
			Name: "KO - Unsufficient balance",
			args: transferArgs{
				pk: func() *ecdsa.PrivateKey {
					key, _ := crypto.GenerateKey()
					return key
				}(),
				To:  crypto.PubkeyToAddress(privKey.PublicKey),
				qty: big.NewInt(1),
			},
			ContractAddr:  *contractAddress,
			ExpectError:   true,
			ExpectedError: "call error on erc20.Transfer(): ERC20InsufficientBalance",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			baseInteractions := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
			if tt.args.pk != nil {
				pk := tt.args.pk
				_, err := baseInteractions.TransferETH(crypto.PubkeyToAddress(pk.PublicKey), big.NewInt(1e18))
				if err != nil {
					t.Fatal(err)
				}

				backend.Commit()
				baseInteractions = base.NewBaseInteractions(backend.Client(), pk, nil, false)
			}
			session, err := erc20.NewIERC20Interactions(
				baseInteractions, tt.ContractAddr, []erc20.BaseERC20Signature{erc20.TransferFrom},
			)
			if err != nil {
				t.Fatal("setting up should not fail")
			}
			_, err = session.TransferTo(tt.args.To, tt.args.qty)
			backend.Commit()
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Contains(t, err.Error(), tt.ExpectedError)
			} else {
				assert.Nil(t, err)
				bal, err := session.BalanceOf(tt.args.To)
				if err != nil {
					t.Fatal("failed to get owner")
				}
				assert.Zero(t, bal.Cmp(tt.args.qty))
			}
		})
	}
}

// Test_GetBalance verifies that the NFT balance is correctly returned for an address.
func Test_GetBalance(t *testing.T) {
	backend, auth, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	token, err := erc20.NewIERC20Interactions(base, *contractAddress, []erc20.BaseERC20Signature{erc20.BalanceOf}, auth)
	assert.Nil(t, err)

	balance, err := token.GetBalance()
	assert.Nil(t, err)
	assert.Equal(t, 0, balance.Cmp(big.NewInt(0).Mul(big.NewInt(100_000_000), big.NewInt(1e18))))
}

// Test_BalanceOf verifies the BalanceOf function for different addresses.
func Test_BalanceOf(t *testing.T) {
	backend, auth, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	token, err := erc20.NewIERC20Interactions(base, *contractAddress, []erc20.BaseERC20Signature{erc20.BalanceOf})
	assert.Nil(t, err)

	testCases := []struct {
		Name           string
		Owner          common.Address
		ExpectedResult int64
		ExpectError    bool
		ExpectedError  string
	}{
		{
			Name:           "OK - non empty balance",
			Owner:          auth.From,
			ExpectedResult: 100_000_000,
		},
		{
			Name:           "OK - empty balance",
			Owner:          common.HexToAddress("1"),
			ExpectedResult: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			balance, err := token.BalanceOf(tt.Owner)
			if tt.ExpectError {
				if err == nil {
					t.Error("expected error but there's none")
					return
				}
				assert.Contains(t, err.Error(), tt.ExpectedError)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, 0, balance.Cmp(big.NewInt(0).Mul(big.NewInt(tt.ExpectedResult), big.NewInt(1e18))))
			}
		})
	}
}

// Test_Approve tests the approval functionality for token transfers.
func Test_Approve(t *testing.T) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	type approveArgs struct {
		To  common.Address
		qty *big.Int
	}

	tests := []struct {
		name          string
		expectError   bool
		args          approveArgs
		errorContains string
	}{
		{
			name: "OK -Successful approval",
			args: approveArgs{
				To:  common.HexToAddress("1"),
				qty: big.NewInt(1),
			},
		},
		{
			name: "NOK - ZeroAddress spender",
			args: approveArgs{
				To:  common.HexToAddress("0"),
				qty: big.NewInt(10),
			},
			expectError:   true,
			errorContains: "call error on erc20.Approve(): ERC20InvalidSpender",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseInteractions := base.NewBaseInteractions(backend.Client(), privKey, nil, false)

			token, err := erc20.NewIERC20Interactions(
				baseInteractions, *contractAddress, []erc20.BaseERC20Signature{erc20.Approve},
			)
			if err != nil {
				t.Fatal(err)
			}

			_, err = token.Approve(tt.args.To, tt.args.qty)
			backend.Commit()

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			} else {
				assert.Nil(t, err)
				approved, err := token.Allowance(baseInteractions.Address, tt.args.To)
				assert.Nil(t, err)
				assert.Equal(t, tt.args.qty.Int64(), approved.Int64())
			}
		})
	}
}

// Test_TokenMetaInfos verifies that the metadata (name, symbol, and URI) for a token is correctly retrieved.
func Test_TokenMetaInfos(t *testing.T) {
	backend, _, contractAddress, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc20MetaData.ABI,
		inferences.Ierc20MetaData.Bin,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	base := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	token, err := erc20.NewIERC20Interactions(base, *contractAddress, []erc20.BaseERC20Signature{erc20.Name, erc20.Symbol})
	assert.Nil(t, err)

	// Test meta infos for token
	tokenInfo, err := token.TokenMetaInfos()
	assert.Nil(t, err)
	assert.Equal(t, "TESTToken", tokenInfo.Name)
	assert.Equal(t, "TT", tokenInfo.Symbol)
}
