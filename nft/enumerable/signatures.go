// Package enumerable provides functions to interact with ERC721 enumerable properties.
package enumerable

import (
	"encoding/hex"

	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/ethereum/go-ethereum/crypto"
)

type IERC721EnumerableSignature nft.BaseNFTSignature

const (
	TokenOfOwnerByIndex IERC721EnumerableSignature = "tokenOfOwnerByIndex(address,uint256)"
	TokenByIndex        IERC721EnumerableSignature = "tokenByIndex(uint256)"
)

func (s IERC721EnumerableSignature) GetHex() string {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(string(s))) // hash.Write never returns an error
	selector := hash.Sum(nil)[:4]
	return hex.EncodeToString(selector)
}
