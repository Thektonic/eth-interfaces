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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IERC20BurnableInteractions wraps interactions with an IERC20Burnable contract, extending basic ERC20 interactions.
type IERC20BurnableInteractions struct {
	*erc20.Interactions
	erc20Burnable *inferences.Ierc20burnable
	callError     func(string, error) error
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

	callError := base.GenCallError("erc20Burnable", ParseError, erc20Burnable.UnpackError)

	return &IERC20BurnableInteractions{baseIERC20, erc20Burnable, callError}, nil
}

// Burn destroys the specified token from the owner's balance.
func (e *IERC20BurnableInteractions) Burn(qty *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		e,
		e,
		e.erc20Burnable.PackBurn(qty),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, e.callError("Burn()", err)
	}

	return tx, nil
}

// BurnFrom is a wrapper for Burn that calls the token's burnFrom function instead.
func (e *IERC20BurnableInteractions) BurnFrom(from common.Address, qty *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		e,
		e,
		e.erc20Burnable.PackBurnFrom(from, qty),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, e.callError("BurnFrom()", err)
	}

	return tx, nil
}

// ParseError parses raw contract errors into human-readable error messages for burnable ERC20 operations.
func ParseError(rawErr any) error {
	switch e := rawErr.(type) {
	case *inferences.Ierc20burnableERC20InsufficientAllowance:
		return fmt.Errorf(
			"ERC20InsufficientAllowance: %s, allowance %s, required: %s",
			e.Spender.Hex(),
			e.Allowance.String(),
			e.Needed.String(),
		)
	case *inferences.Ierc20burnableERC20InvalidSpender:
		return fmt.Errorf("ERC20InvalidSpender: %s", e.Spender.Hex())
	case *inferences.Ierc20burnableERC20InsufficientBalance:
		return fmt.Errorf("ERC20InsufficientBalance: %s, required: %s", e.Balance.String(), e.Needed.String())
	case *inferences.Ierc20burnableERC20InvalidSender:
		return fmt.Errorf("ERC20InvalidSender: %s", e.Sender.Hex())
	case *inferences.Ierc20burnableERC20InvalidReceiver:
		return fmt.Errorf("ERC20InvalidReceiver: %s", e.Receiver.Hex())
	case *inferences.Ierc20burnableERC20InvalidApprover:
		return fmt.Errorf("ERC20InvalidApprover: %s", e.Approver.Hex())
	default:
		return nil
	}
}
