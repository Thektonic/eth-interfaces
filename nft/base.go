package nft

// Package nft provides base functionality for interacting with NFTs using the IERC721 standard.

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/contractextension"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721Complete"
	"github.com/Thektonic/eth-interfaces/models"
	"github.com/Thektonic/eth-interfaces/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC721Interactions wraps NFT interactions using an underlying base interaction and an ERC721A session.

type ERC721Interactions struct {
	*base.Interactions[*ERC721Complete.ERC721CompleteSession]
}

type INFT interface {
	GetAddress() common.Address
	CheckSignatures(address common.Address, signatures []utils.Signature) error
	GetSession() ERC721Complete.ERC721CompleteSession
	GetBaseInteractions() *base.BaseInteractions
}

// NewERC721Interactions creates a new instance of ERC721Interactions from a base interaction interface and an NFT contract address.
func NewERC721Interactions(
	baseInteractions *base.BaseInteractions,
	address common.Address,
	signatures []BaseNFTSignature,
	transactOps ...*bind.TransactOpts,
) (*ERC721Interactions, error) {

	var converted []utils.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	err := baseInteractions.CheckSignatures(address, converted)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("CheckSignatures", err)
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

	erc721Complete, err := ERC721Complete.NewERC721Complete(address, baseInteractions.Client)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("NewERC721Interactions", err)
	}
	erc721ASession := ERC721Complete.ERC721CompleteSession{
		Contract:     erc721Complete,
		CallOpts:     bind.CallOpts{Pending: true, From: baseInteractions.Address},
		TransactOpts: *txOpts,
	}

	callError := func(field string, err error) *base.CallError {
		return baseInteractions.WrapCallError(ERC721Complete.ERC721CompleteABI, field, err)
	}

	erc721Interactions := &ERC721Interactions{
		&base.Interactions[*ERC721Complete.ERC721CompleteSession]{
			BaseInteractions: baseInteractions,
			Session:          &erc721ASession,
			Address:          address,
			CallError:        callError,
		},
	}

	if err := contractextension.SimulateCall(baseInteractions.Ctx, ERC721Complete.ERC721CompleteABI, "name", erc721Interactions); err != nil {
		return nil, err
	}

	return erc721Interactions, nil
}

// GetSession returns the current session used for NFT interactions.
func (d *ERC721Interactions) GetSession() ERC721Complete.ERC721CompleteSession {
	return *d.Session
}

// GetBalance retrieves the balance of NFTs for the associated address.
func (d *ERC721Interactions) GetBalance() (*big.Int, error) {
	balance, err := d.Session.BalanceOf(d.Address)
	if err != nil {
		return nil, d.CallError("nft.BalanceOf()", err)
	}
	return balance, nil
}

// TransferTo transfers a specific token to another address after verifying ownership.
func (d *ERC721Interactions) TransferTo(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := d.Session.TransferFrom(d.Address, to, tokenID)
	if err != nil {
		return nil, d.CallError("nft.TransferFrom()", err)
	}
	return tx, nil
}

// TransferFirstOwnedTo transfers the first token owned by the signer to the specified address.
func (d *ERC721Interactions) TransferFirstOwnedTo(to common.Address) (*types.Transaction, error) {
	maxSupply, err := d.TotalSupply()
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: %w", err)
	}

	for idx := range maxSupply.Int64() {
		tokenID := big.NewInt(idx)
		tx, err := d.TransferTo(to, tokenID)
		if err != nil {
			if strings.Contains(err.Error(), utils.ErrZeroAddress.Error()) {
				return nil, err
			}
			continue
		}
		return tx, nil
	}

	return nil, errors.New("no nft found from signer")
}

// TotalSupply returns the total number of NFTs minted.
func (d *ERC721Interactions) TotalSupply() (*big.Int, error) {
	supply, err := d.Session.TotalSupply()
	if err != nil {
		return nil, d.CallError("nft.TotalSupply()", err)
	}
	return supply, nil
}

// BalanceOf retrieves the NFT balance for a given owner.
func (d *ERC721Interactions) BalanceOf(owner common.Address) (*big.Int, error) {
	balance, err := d.Session.BalanceOf(owner)
	if err != nil {
		return nil, d.CallError("nft.BalanceOf()", err)
	}
	return balance, nil
}

// OwnerOf retrieves the owner of a specific token.
func (d *ERC721Interactions) OwnerOf(tokenID *big.Int) (common.Address, error) {
	owner, err := d.Session.OwnerOf(tokenID)
	if err != nil {
		return common.Address{}, d.CallError("nft.OwnerOf()", err)
	}
	return owner, nil
}

// Approve approves an address to transfer a specific token.
func (d *ERC721Interactions) Approve(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := d.Session.Approve(to, tokenID)
	if err != nil {
		return nil, d.CallError("nft.Approve()", err)
	}
	return tx, nil
}

// TokenMetaInfos retrieves metadata about the specified token such as name, symbol, and URI.
func (d *ERC721Interactions) TokenMetaInfos(tokenID *big.Int) (*models.TokenMeta, error) {
	name, err := d.Name()
	if err != nil {
		return nil, err
	}
	symbol, err := d.Symbol()
	if err != nil {
		return &models.TokenMeta{Name: name}, err
	}

	uri, err := d.TokenURI(tokenID)
	if err != nil {
		return &models.TokenMeta{Name: name, Symbol: symbol}, err
	}

	return &models.TokenMeta{Name: name, Symbol: symbol, URI: uri}, nil
}

// Name returns the name of the NFT.
func (d *ERC721Interactions) Name() (string, error) {
	name, err := d.Session.Name()
	if err != nil {
		return "", d.CallError("nft.Name()", err)
	}
	return name, nil
}

// Symbol returns the symbol of the NFT.
func (d *ERC721Interactions) Symbol() (string, error) {
	symbol, err := d.Session.Symbol()
	if err != nil {
		return "", d.CallError("nft.Symbol()", err)
	}
	return symbol, nil
}

// TokenURI returns the URI of the NFT.
func (d *ERC721Interactions) TokenURI(tokenID *big.Int) (string, error) {
	uri, err := d.Session.TokenURI(tokenID)
	if err != nil {
		return "", d.CallError("nft.TokenURI()", err)
	}
	return uri, nil
}

// GetApproved returns the approved address for a specific token.
func (d *ERC721Interactions) GetApproved(tokenID *big.Int) (common.Address, error) {
	approved, err := d.Session.GetApproved(tokenID)
	if err != nil {
		return common.Address{}, d.CallError("nft.GetApproved()", err)
	}
	return approved, nil
}
