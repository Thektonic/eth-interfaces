// Package nft provides base functionality for interacting with NFTs using the IERC721 standard.
package nft

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

type BaseNFTSignature string

const (
	Name             BaseNFTSignature = "name()"
	Symbol           BaseNFTSignature = "symbol()"
	BalanceOf        BaseNFTSignature = "balanceOf(address)"
	TotalSupply      BaseNFTSignature = "totalSupply()"
	OwnerOf          BaseNFTSignature = "ownerOf(uint256)"
	TokenURI         BaseNFTSignature = "tokenURI(uint256)" // #nosec G101
	Approve          BaseNFTSignature = "approve(address,uint256)"
	GetApproved      BaseNFTSignature = "getApproved(uint256)"
	TransferFrom     BaseNFTSignature = "transferFrom(address,address,uint256)"
	SafeTransferFrom BaseNFTSignature = "safeTransferFrom(address,address,uint256)"
)

func (s BaseNFTSignature) GetHex() string {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(s)) // hash.Write never returns an error
	selector := hash.Sum(nil)[:4]
	return hex.EncodeToString(selector)
}
