// Package erc20 provides base functionality for interacting with ERC20 tokens using the IERC20 standard.
package erc20

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// BaseERC20Signature represents function signatures for basic ERC20 token operations
type BaseERC20Signature string

const (
	// Name represents the name function signature
	Name BaseERC20Signature = "name()"
	// Symbol represents the symbol function signature
	Symbol BaseERC20Signature = "symbol()"
	// Decimals represents the decimals function signature
	Decimals BaseERC20Signature = "decimals()"
	// BalanceOf represents the balanceOf function signature
	BalanceOf BaseERC20Signature = "balanceOf(address)"
	// TotalSupply represents the totalSupply function signature
	TotalSupply BaseERC20Signature = "totalSupply()"
	// TokenURI represents the tokenURI function signature
	TokenURI BaseERC20Signature = "tokenURI(uint256)" // #nosec G101
	// Approve represents the approve function signature
	Approve BaseERC20Signature = "approve(address,uint256)"
	// TransferFrom represents the transferFrom function signature
	TransferFrom BaseERC20Signature = "transferFrom(address,address,uint256)"
	// SafeTransferFrom represents the safeTransferFrom function signature
	SafeTransferFrom BaseERC20Signature = "safeTransferFrom(address,address,uint256)"
)

// GetHex returns the hex representation of the function signature
func (s BaseERC20Signature) GetHex() string {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(s)) // hash.Write never returns an error
	selector := hash.Sum(nil)[:4]
	return hex.EncodeToString(selector)
}
