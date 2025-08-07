// Package enumerable provides functions to interact with ERC721 enumerable properties.
package enumerable

import (
	"encoding/hex"

	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/ethereum/go-ethereum/crypto"
)

// IERC721EnumerableSignature represents function signatures for ERC721 enumerable operations
type IERC721EnumerableSignature nft.BaseNFTSignature

const (
	// TokenOfOwnerByIndex represents the tokenOfOwnerByIndex function signature
	TokenOfOwnerByIndex IERC721EnumerableSignature = "tokenOfOwnerByIndex(address,uint256)"
	// TokenByIndex represents the tokenByIndex function signature
	TokenByIndex IERC721EnumerableSignature = "tokenByIndex(uint256)" // #nosec G101
)

// GetHex returns the hex representation of the function signature
func (s IERC721EnumerableSignature) GetHex() string {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(string(s))) // hash.Write never returns an error
	selector := hash.Sum(nil)
	return hex.EncodeToString(selector)
}

func (s IERC721EnumerableSignature) String() string {
	return string(s)
}

func (s IERC721EnumerableSignature) GetSelector() []byte {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(string(s))) // hash.Write never returns an error
	return hash.Sum(nil)
}
