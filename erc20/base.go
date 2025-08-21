// Package erc20 provides base functionality for interacting with ERC20 tokens using the IERC20 standard.
package erc20

import (
	"fmt"
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/models"
	"github.com/Thektonic/eth-interfaces/transaction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IERC20 interface defines the functions for NFT interactions.

// IERC20AInteractions wraps NFT interactions using an underlying base interaction and an ERC20 session.

// Interactions provides methods for interacting with ERC20 token contracts

type session struct {
	erc20    *inferences.Ierc20
	callOpts *bind.CallOpts
	instance *bind.BoundContract
}

func (s *session) CallOpts() *bind.CallOpts {
	return s.callOpts
}
func (s *session) Instance() *bind.BoundContract {
	return s.instance
}

// Interactions provides methods for interacting with ERC20 token contracts.
type Interactions struct {
	*base.Interactions
	*session
	erc20Address common.Address
	callError    func(string, error) *base.CallError
}

// NewIERC20Interactions creates a new instance of IERC20AInteractions from a base interaction
// interface and an NFT contract address.
func NewIERC20Interactions(
	baseInteractions *base.Interactions,
	address common.Address,
	signatures []BaseERC20Signature,
	transactOps ...*bind.TransactOpts,
) (*Interactions, error) {
	var converted []hex.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	err := baseInteractions.CheckSignatures(address, converted)
	if err != nil {
		return nil, err
	}

	ierc20 := inferences.NewIerc20()

	ierc20Session := &session{
		erc20:    ierc20,
		callOpts: &bind.CallOpts{Pending: true, From: baseInteractions.Address},
		instance: ierc20.Instance(baseInteractions.Client, address),
	}

	callError := func(field string, err error) *base.CallError {
		return (baseInteractions.WrapCallError(inferences.Ierc20MetaData.ABI, field, err))
	}

	ierc20Asession := &Interactions{
		baseInteractions,
		ierc20Session,
		address,
		callError,
	}

	if len(transactOps) > 0 {
		if transactOps[0] == nil {
			return nil, fmt.Errorf("transactOpts cannot be nil")
		}
		ierc20Asession.TxOptsFn = func() (*bind.TransactOpts, error) {
			return transactOps[0], nil
		}
	}

	return ierc20Asession, nil
}

// GetAddress returns the ERC20 contract address.
func (d *Interactions) GetAddress() common.Address {
	return d.erc20Address
}

// GetSession returns the current session used for NFT interactions.
func (d *Interactions) GetSession() transaction.Session {
	return d.session
}

// GetBalance retrieves the balance of NFTs for the associated address.
func (d *Interactions) GetBalance() (*big.Int, error) {
	balance, err := transaction.Call(
		d.session,
		d.erc20.PackBalanceOf(d.Address),
		d.erc20.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("erc20.BalanceOf()", err)
	}
	return balance, nil
}

// TransferTo transfers a specific token to another address after verifying ownership.
func (d *Interactions) TransferTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.erc20.PackTransfer(to, amount),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("erc20.Transfer()", err)
	}
	return tx, nil
}

// Decimals returns the number of decimals used to get its user representation.
func (d *Interactions) Decimals() (uint8, error) {
	decimals, err := transaction.Call(
		d.session,
		d.erc20.PackDecimals(),
		d.erc20.UnpackDecimals,
	)

	if err != nil {
		return 0, d.callError("erc20.Decimals()", err)
	}
	return decimals, nil
}

// TotalSupply returns the total number of NFTs minted.
func (d *Interactions) TotalSupply() (*big.Int, error) {
	supply, err := transaction.Call(
		d.session,
		d.erc20.PackTotalSupply(),
		d.erc20.UnpackTotalSupply,
	)

	if err != nil {
		return nil, d.callError("erc20.TotalSupply()", err)
	}
	return supply, nil
}

// BalanceOf retrieves the NFT balance for a given owner.
func (d *Interactions) BalanceOf(owner common.Address) (*big.Int, error) {
	balance, err := transaction.Call(
		d.session,
		d.erc20.PackBalanceOf(owner),
		d.erc20.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("erc20.BalanceOf()", err)
	}
	return balance, nil
}

// Approve approves an address to transfer a specific token.
func (d *Interactions) Approve(to common.Address, allowance *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.erc20.PackApprove(to, allowance), transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("erc20.Approve()", err)
	}
	return tx, nil
}

// TokenMetaInfos retrieves metadata about the specified token such as name, symbol, and URI.
func (d *Interactions) TokenMetaInfos() (*models.TokenMeta, error) {
	name, err := d.Name()
	if err != nil {
		return nil, err
	}
	symbol, err := d.Symbol()
	if err != nil {
		return &models.TokenMeta{Name: name}, err
	}
	return &models.TokenMeta{Name: name, Symbol: symbol}, nil
}

// Name returns the name of the NFT.
func (d *Interactions) Name() (string, error) {
	name, err := transaction.Call(
		d.session,
		d.erc20.PackName(),
		d.erc20.UnpackName,
	)

	if err != nil {
		return "", d.callError("erc20.Name()", err)
	}

	return name, nil
}

// Symbol returns the symbol of the NFT.
func (d *Interactions) Symbol() (string, error) {
	symbol, err := transaction.Call(
		d.session,
		d.erc20.PackSymbol(),
		d.erc20.UnpackSymbol,
	)

	if err != nil {
		return "", d.callError("erc20.Symbol()", err)
	}
	return symbol, nil
}

// Allowance returns the amount of tokens that spender is allowed to spend on behalf of owner
func (d *Interactions) Allowance(owner, spender common.Address) (*big.Int, error) {
	allowance, err := transaction.Call(
		d.session,
		d.erc20.PackAllowance(owner, spender),
		d.erc20.UnpackAllowance,
	)

	if err != nil {
		return nil, d.callError("erc20.Allowance()", err)
	}

	return allowance, nil
}
