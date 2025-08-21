// Package enumerable provides functions to interact with ERC721 enumerable properties.
package enumerable

import (
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/Thektonic/eth-interfaces/transaction"
	"github.com/ethereum/go-ethereum/common"
)

// ERC721EnumerableInteractions wraps interactions with an ERC721Enumerable contract, extending basic NFT interactions.
type ERC721EnumerableInteractions struct {
	*nft.ERC721Interactions
	ierc721Enumerable *inferences.Ierc721
	callError         func(string, error) *base.CallError
}

// NewERC721EnumerableInteractions creates a new enumerable interaction instance
// using the provided base NFT interactions.
func NewERC721EnumerableInteractions(
	baseIERC721 *nft.ERC721Interactions,
	signatures []IERC721EnumerableSignature,
) (*ERC721EnumerableInteractions, error) {
	var converted []hex.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}
	err := baseIERC721.CheckSignatures(baseIERC721.GetAddress(), converted)
	if err != nil {
		return nil, customerrors.WrapInterfacingError("erc721Enumerable", err)
	}

	ierc721Enumerable := inferences.NewIerc721()

	callError := func(field string, err error) *base.CallError {
		return baseIERC721.WrapCallError(inferences.Ierc721MetaData.ABI, field, err)
	}

	return &ERC721EnumerableInteractions{baseIERC721, ierc721Enumerable, callError}, nil
}

// GetAddressOwnedTokens returns a slice of token IDs owned by the specified address.
func (e *ERC721EnumerableInteractions) GetAddressOwnedTokens(to common.Address) ([]*big.Int, error) {
	balance, err := e.BalanceOf(to)
	if err != nil {
		return nil, err
	}
	tokenIDs := []*big.Int{}
	for i := range balance.Int64() {
		tokenID, err := e.TokenOfOwnerByIndex(to, big.NewInt(i))
		if err != nil {
			return nil, e.callError("nft.TokenOfOwnerByIndex()", err)
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	return tokenIDs, nil
}

// GetAllTokenIDs returns all token IDs available in the contract.
func (e *ERC721EnumerableInteractions) GetAllTokenIDs() ([]*big.Int, error) {
	supply, err := e.TotalSupply()
	if err != nil {
		return nil, err
	}
	tokenIDs := []*big.Int{}
	for i := range supply.Int64() {
		tokenID, err := e.TokenByIndex(big.NewInt(i))
		if err != nil {
			return nil, e.callError("nft.TokenByIndex()", err)
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	return tokenIDs, nil
}

// TokenOfOwnerByIndex returns the token ID belonging to a specified address at a given index.
func (e *ERC721EnumerableInteractions) TokenOfOwnerByIndex(to common.Address, index *big.Int) (*big.Int, error) {
	tokenID, err := transaction.Call(
		e.GetSession(),
		e.ierc721Enumerable.PackTokenOfOwnerByIndex(to, index),
		e.ierc721Enumerable.UnpackTokenOfOwnerByIndex,
	)

	if err != nil {
		return nil, e.callError("nft.TokenOfOwnerByIndex()", err)
	}
	return tokenID, nil
}

// TokenByIndex returns the token ID at a specific index in the contract.
func (e *ERC721EnumerableInteractions) TokenByIndex(index *big.Int) (*big.Int, error) {
	tokenID, err := transaction.Call(
		e.GetSession(),
		e.ierc721Enumerable.PackTokenByIndex(index),
		e.ierc721Enumerable.UnpackTokenByIndex,
	)

	if err != nil {
		return nil, e.callError("nft.TokenByIndex()", err)
	}
	return tokenID, nil
}
