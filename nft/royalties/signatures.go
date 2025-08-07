// Package royalties provides functions to interact with ERC721 royalty properties.
package royalties

import (
	"encoding/hex"

	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/ethereum/go-ethereum/crypto"
)

// IERC721RoyaltiesSignature represents function signatures for ERC721 royalties operations
type IERC721RoyaltiesSignature nft.BaseNFTSignature

const (
	// RoyaltyInfo represents the royaltyInfo function signature
	RoyaltyInfo IERC721RoyaltiesSignature = "royaltyInfo(uint256,uint256)"
)

// computeHash returns the Keccak256 hash of the function signature
func (s IERC721RoyaltiesSignature) computeHash() []byte {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(string(s))) // hash.Write never returns an error
	return hash.Sum(nil)
}

// GetHex returns the hex representation of the function signature
func (s IERC721RoyaltiesSignature) GetHex() string {
	return hex.EncodeToString(s.computeHash())
}

func (s IERC721RoyaltiesSignature) String() string {
	return string(s)
}

// GetSelector returns the Keccak256 hash selector for the ERC721 royalties signature
func (s IERC721RoyaltiesSignature) GetSelector() []byte {
	return s.computeHash()[:4]
}
