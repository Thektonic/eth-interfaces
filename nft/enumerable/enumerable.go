package enumerable

// Package enumerable provides functions to interact with ERC721 enumerable properties.

import (
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721Complete"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/Thektonic/eth-interfaces/utils"
	"github.com/ethereum/go-ethereum/common"
)

// ERC721EnumerableInteractions wraps interactions with an ERC721Enumerable contract, extending basic NFT interactions.
type ERC721EnumerableInteractions struct {
	*base.Interactions[*ERC721Complete.ERC721CompleteSession]
}

// NewERC721EnumerableInteractions creates a new enumerable interaction instance using the provided base NFT interactions.
func NewERC721EnumerableInteractions(baseIERC721 nft.INFT, signatures []IERC721EnumerableSignature) (*ERC721EnumerableInteractions, error) {
	ierc721Enumerable, err := ERC721Complete.NewERC721Complete(baseIERC721.GetAddress(), baseIERC721.GetBaseInteractions().Client)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("erc721Enumerable", err)
	}
	session := ERC721Complete.ERC721CompleteSession{
		Contract:     ierc721Enumerable,
		CallOpts:     baseIERC721.GetSession().CallOpts,
		TransactOpts: baseIERC721.GetSession().TransactOpts,
	}

	var converted []utils.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	callError := func(field string, err error) *base.CallError {
		return baseIERC721.GetBaseInteractions().WrapCallError(ERC721Complete.ERC721CompleteABI, field, err)
	}

	err = baseIERC721.CheckSignatures(baseIERC721.GetAddress(), converted)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("erc721Enumerable", err)
	}

	return &ERC721EnumerableInteractions{
		&base.Interactions[*ERC721Complete.ERC721CompleteSession]{
			BaseInteractions: baseIERC721.GetBaseInteractions(),
			Session:          &session,
			Address:          baseIERC721.GetAddress(),
			CallError:        callError,
		},
	}, nil
}

// GetAddressOwnedTokens returns a slice of token IDs owned by the specified address.
func (e *ERC721EnumerableInteractions) GetAddressOwnedTokens(to common.Address) ([]*big.Int, error) {
	balance, err := e.Session.BalanceOf(to)
	if err != nil {
		return nil, err
	}
	tokenIDs := []*big.Int{}
	for i := range balance.Int64() {
		tokenID, err := e.TokenOfOwnerByIndex(to, big.NewInt(i))
		if err != nil {
			return nil, e.CallError("nft.TokenOfOwnerByIndex()", err)
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	return tokenIDs, nil
}

// GetAllTokenIDs returns all token IDs available in the contract.
func (e *ERC721EnumerableInteractions) GetAllTokenIDs() ([]*big.Int, error) {
	supply, err := e.Session.TotalSupply()
	if err != nil {
		return nil, err
	}
	tokenIDs := []*big.Int{}
	for i := range supply.Int64() {
		tokenID, err := e.TokenByIndex(big.NewInt(i))
		if err != nil {
			return nil, e.CallError("nft.TokenByIndex()", err)
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	return tokenIDs, nil
}

// TokenOfOwnerByIndex returns the token ID belonging to a specified address at a given index.
func (e *ERC721EnumerableInteractions) TokenOfOwnerByIndex(to common.Address, index *big.Int) (*big.Int, error) {
	tokenID, err := e.Session.TokenOfOwnerByIndex(to, index)
	if err != nil {
		return nil, e.CallError("nft.TokenOfOwnerByIndex()", err)
	}
	return tokenID, nil
}

// TokenByIndex returns the token ID at a specific index in the contract.
func (e *ERC721EnumerableInteractions) TokenByIndex(index *big.Int) (*big.Int, error) {
	tokenID, err := e.Session.TokenByIndex(index)
	if err != nil {
		return nil, e.CallError("nft.TokenByIndex()", err)
	}
	return tokenID, nil
}
