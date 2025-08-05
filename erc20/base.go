// Package erc20 provides base functionality for interacting with ERC20 tokens using the IERC20 standard.
package erc20

import (
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/contractextension"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences/ERC20Burnable"
	"github.com/Thektonic/eth-interfaces/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IERC20 interface defines the functions for NFT interactions.

// IERC20AInteractions wraps NFT interactions using an underlying base interaction and an ERC20 session.

// Interactions provides methods for interacting with ERC20 token contracts
type Interactions struct {
	*base.Interactions
	ierc20Session *ERC20Burnable.ERC20BurnableSession
	erc20Address  common.Address
	callError     func(string, error) *base.CallError
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

	var txOpts *bind.TransactOpts
	if len(transactOps) == 0 {
		txOpts, err = baseInteractions.BaseTxSetup()
		if err != nil {
			return nil, customerrors.WrapinterfacingError("BaseTxSetup", err)
		}
	} else {
		txOpts = transactOps[0]
	}

	ierc20, err := ERC20Burnable.NewERC20Burnable(address, baseInteractions.Client)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("NewIERC20", err)
	}
	ierc20Session := ERC20Burnable.ERC20BurnableSession{
		Contract:     ierc20,
		CallOpts:     bind.CallOpts{Pending: true, From: baseInteractions.Address},
		TransactOpts: *txOpts,
	}

	callError := func(field string, err error) *base.CallError {
		return (baseInteractions.WrapCallError(ERC20Burnable.ERC20BurnableABI, field, err))
	}

	ierc20Asession := &Interactions{baseInteractions,
		&ierc20Session,
		address,
		callError,
	}

	if err := contractextension.SimulateCall(
		baseInteractions.Ctx, ERC20Burnable.ERC20BurnableABI, "name", ierc20Asession,
	); err != nil {
		return nil, err
	}

	return ierc20Asession, nil
}

// GetAddress returns the ERC20 contract address.
func (d *Interactions) GetAddress() common.Address {
	return d.erc20Address
}

// GetSession returns the current session used for NFT interactions.
func (d *Interactions) GetSession() ERC20Burnable.ERC20BurnableSession {
	return *d.ierc20Session
}

// GetBalance retrieves the balance of NFTs for the associated address.
func (d *Interactions) GetBalance() (*big.Int, error) {
	balance, err := d.ierc20Session.BalanceOf(d.Address)
	if err != nil {
		return nil, d.callError("erc20.BalanceOf()", err)
	}
	return balance, nil
}

// TransferTo transfers a specific token to another address after verifying ownership.
func (d *Interactions) TransferTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := d.ierc20Session.Transfer(to, amount)
	if err != nil {
		return nil, d.callError("erc20.Transfer()", err)
	}
	return tx, nil
}

// Decimals returns the number of decimals used to get its user representation.
func (d *Interactions) Decimals() (uint8, error) {
	decimals, err := d.ierc20Session.Decimals()
	if err != nil {
		return 0, d.callError("erc20.Decimals()", err)
	}
	return decimals, nil
}

// TotalSupply returns the total number of NFTs minted.
func (d *Interactions) TotalSupply() (*big.Int, error) {
	supply, err := d.ierc20Session.TotalSupply()
	if err != nil {
		return nil, d.callError("erc20.TotalSupply()", err)
	}
	return supply, nil
}

// BalanceOf retrieves the NFT balance for a given owner.
func (d *Interactions) BalanceOf(owner common.Address) (*big.Int, error) {
	balance, err := d.ierc20Session.BalanceOf(owner)
	if err != nil {
		return nil, d.callError("erc20.BalanceOf()", err)
	}
	return balance, nil
}

// Approve approves an address to transfer a specific token.
func (d *Interactions) Approve(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := d.ierc20Session.Approve(to, tokenID)
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
	name, err := d.ierc20Session.Name()
	if err != nil {
		return "", d.callError("erc20.Name()", err)
	}
	return name, nil
}

// Symbol returns the symbol of the NFT.
func (d *Interactions) Symbol() (string, error) {
	symbol, err := d.ierc20Session.Symbol()
	if err != nil {
		return "", d.callError("erc20.Symbol()", err)
	}
	return symbol, nil
}

// Allowance returns the amount of tokens that spender is allowed to spend on behalf of owner
func (d *Interactions) Allowance(owner, spender common.Address) (*big.Int, error) {
	allowance, err := d.ierc20Session.Allowance(owner, spender)
	if err != nil {
		return nil, d.callError("erc20.Allowance()", err)
	}
	return allowance, nil
}
