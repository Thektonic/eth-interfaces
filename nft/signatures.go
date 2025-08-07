// Package nft provides base functionality for interacting with NFTs using the IERC721 standard.
package nft

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// BaseNFTSignature represents function signatures for basic NFT operations
type BaseNFTSignature string

const (
	// Name represents the name function signature
	Name BaseNFTSignature = "name()"
	// Symbol represents the symbol function signature
	Symbol BaseNFTSignature = "symbol()"
	// BalanceOf represents the balanceOf function signature
	BalanceOf BaseNFTSignature = "balanceOf(address)"
	// TotalSupply represents the totalSupply function signature
	TotalSupply BaseNFTSignature = "totalSupply()"
	// OwnerOf represents the ownerOf function signature
	OwnerOf BaseNFTSignature = "ownerOf(uint256)"
	// TokenURI represents the tokenURI function signature
	TokenURI BaseNFTSignature = "tokenURI(uint256)" // #nosec G101
	// Approve represents the approve function signature
	Approve BaseNFTSignature = "approve(address,uint256)"
	// GetApproved represents the getApproved function signature
	GetApproved BaseNFTSignature = "getApproved(uint256)"
	// TransferFrom represents the transferFrom function signature
	TransferFrom BaseNFTSignature = "transferFrom(address,address,uint256)"
	// SafeTransferFrom represents the safeTransferFrom function signature
	SafeTransferFrom BaseNFTSignature = "safeTransferFrom(address,address,uint256)"
)

// computeHash returns the Keccak256 hash of the function signature
func (s BaseNFTSignature) computeHash() []byte {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(s)) // hash.Write never returns an error
	return hash.Sum(nil)
}

// GetHex returns the hex representation of the function signature
func (s BaseNFTSignature) GetHex() string {
	return hex.EncodeToString(s.computeHash())
}

func (s BaseNFTSignature) String() string {
	return string(s)
}

// GetSelector returns the Keccak256 hash selector for the base NFT signature
func (s BaseNFTSignature) GetSelector() []byte {
	return s.computeHash()[:4]
}
