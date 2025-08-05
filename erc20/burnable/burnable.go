package burnable

// Package enumerable provides functions to interact with ERC721 enumerable properties.

import (
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/erc20"
	"github.com/Thektonic/eth-interfaces/inferences/ERC20Burnable"
	"github.com/Thektonic/eth-interfaces/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IERC20Burnable wraps interactions with an IERC20Burnable contract, extending basic NFT interactions.
type IERC20BurnableInteractions struct {
	*base.Interactions[*ERC20Burnable.ERC20BurnableSession]
}

// NewIERC20Burnable creates a new enumerable interaction instance using the provided base NFT interactions.
func NewIERC20Burnable(baseIERC20 erc20.IERC20, signatures []ERC20BurnableSignatures) (*IERC20BurnableInteractions, error) {
	erc20Burnable, err := ERC20Burnable.NewERC20Burnable(baseIERC20.GetAddress(), baseIERC20.GetBaseInteractions().Client)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("ierc20Burnable", err)
	}
	session := ERC20Burnable.ERC20BurnableSession{
		Contract:     erc20Burnable,
		CallOpts:     baseIERC20.GetSession().CallOpts,
		TransactOpts: baseIERC20.GetSession().TransactOpts,
	}

	var converted []utils.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	err = baseIERC20.CheckSignatures(baseIERC20.GetAddress(), converted)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("ierc20Burnable", err)
	}

	callError := func(field string, err error) *base.CallError {
		return baseIERC20.GetBaseInteractions().WrapCallError(ERC20Burnable.ERC20BurnableABI, field, err)
	}

	return &IERC20BurnableInteractions{
		&base.Interactions[*ERC20Burnable.ERC20BurnableSession]{
			BaseInteractions: baseIERC20.GetBaseInteractions(),
			Session:          &session,
			Address:          baseIERC20.GetAddress(),
			CallError:        callError,
		},
	}, nil
}

// Burn destroys the specified token from the owner's balance.
func (e *IERC20BurnableInteractions) Burn(qty *big.Int) (*types.Transaction, error) {
	tx, err := e.Session.Burn(qty)
	if err != nil {
		return nil, e.CallError("erc20.Burn()", err)
	}
	return tx, nil
}

// BurnFrom is a wrapper for Burn that calls the token's burnFrom function instead.
func (e *IERC20BurnableInteractions) BurnFrom(from common.Address, qty *big.Int) (*types.Transaction, error) {
	tx, err := e.Session.BurnFrom(from, qty)
	if err != nil {
		return nil, e.CallError("nft.BurnFrom()", err)
	}
	return tx, nil
}
