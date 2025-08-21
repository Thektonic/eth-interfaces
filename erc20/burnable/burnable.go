// Package burnable provides functions to interact with ERC20 burnable properties.
package burnable

import (
	"fmt"
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/erc20"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/transaction"
	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IERC20BurnableInteractions wraps interactions with an IERC20Burnable contract, extending basic ERC20 interactions.
type IERC20BurnableInteractions struct {
	*erc20.Interactions
	erc20Burnable *inferences.Ierc20burnable
	callError     func(string, error) *base.CallError
}

// NewIERC20Burnable creates a new enumerable interaction instance using the provided base NFT interactions.
func NewIERC20Burnable(
	baseIERC20 *erc20.Interactions,
	signatures []ERC20BurnableSignatures,
) (*IERC20BurnableInteractions, error) {
	var converted []hex.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	err := baseIERC20.CheckSignatures(baseIERC20.GetAddress(), converted)
	if err != nil {
		return nil, customerrors.WrapInterfacingError("ierc20Burnable", err)
	}

	erc20Burnable := inferences.NewIerc20burnable()

	callError := func(field string, err error) *base.CallError {
		return baseIERC20.WrapCallError(inferences.Ierc20burnableMetaData.ABI, field, err)
	}

	return &IERC20BurnableInteractions{baseIERC20, erc20Burnable, callError}, nil
}

// Burn destroys the specified token from the owner's balance.
func (e *IERC20BurnableInteractions) Burn(qty *big.Int) (*types.Transaction, error) {
	if e.Safe() {
		_, err := transaction.Call(e, e.erc20Burnable.PackBurn(qty), transaction.DefaultUnpacker)
		if err != nil {
			fmt.Println(err.Error())
			return nil, e.callError("erc20.Approve()", err)
		}
	}

	txOpts, err := e.BaseTxSetup()
	if err != nil {
		return nil, e.callError("BaseTxSetup()", err)
	}
	tx, err := bind2.Transact(e.Instance(), txOpts, e.erc20Burnable.PackBurn(qty))
	if err != nil {
		return nil, e.callError("erc20.Burn()", err)
	}

	return tx, nil
}

// BurnFrom is a wrapper for Burn that calls the token's burnFrom function instead.
func (e *IERC20BurnableInteractions) BurnFrom(from common.Address, qty *big.Int) (*types.Transaction, error) {
	if e.Safe() {
		_, err := transaction.Call(e.Interactions, e.erc20Burnable.PackBurnFrom(from, qty), transaction.DefaultUnpacker)
		if err != nil {
			fmt.Println(err.Error())
			return nil, e.callError("erc20.Approve()", err)
		}
	}

	txOpts, err := e.BaseTxSetup()
	if err != nil {
		return nil, e.callError("BaseTxSetup()", err)
	}
	tx, err := bind2.Transact(e.Instance(), txOpts, e.erc20Burnable.PackBurnFrom(from, qty))
	if err != nil {
		return nil, e.callError("erc20.BurnFrom()", err)
	}

	return tx, nil
}
