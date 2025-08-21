// Package nft provides base functionality for interacting with NFTs using the IERC721 standard.
package nft

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/models"
	"github.com/Thektonic/eth-interfaces/transaction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type session struct {
	erc271   *inferences.Ierc721
	callOpts *bind.CallOpts
	instance *bind.BoundContract
}

func (s *session) CallOpts() *bind.CallOpts {
	return s.callOpts
}
func (s *session) Instance() *bind.BoundContract {
	return s.instance
}

// ERC721Interactions provides methods for interacting with ERC721 NFT contracts
type ERC721Interactions struct {
	*base.Interactions
	*session
	nftAddress common.Address
	callError  func(string, error) *base.CallError
}

// NewERC721Interactions creates a new instance of ERC721Interactions from a base interaction
// interface and an NFT contract address.
func NewERC721Interactions(
	baseInteractions *base.Interactions,
	address common.Address,
	signatures []BaseNFTSignature,
	transactOps ...*bind.TransactOpts,
) (*ERC721Interactions, error) {
	var converted []hex.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	if err := baseInteractions.CheckSignatures(address, converted); err != nil {
		return nil, customerrors.WrapinterfacingError("CheckSignatures", err)
	}

	erc721Complete := inferences.NewIerc721()

	erc721ASession := session{
		erc271:   erc721Complete,
		callOpts: &bind.CallOpts{Pending: true, From: baseInteractions.Address},
		instance: erc721Complete.Instance(baseInteractions.Client, address),
	}

	callError := func(field string, err error) *base.CallError {
		return baseInteractions.WrapCallError(inferences.Ierc721MetaData.ABI, field, err)
	}

	erc721Interactions := &ERC721Interactions{baseInteractions,
		&erc721ASession,
		address,
		callError,
	}

	if len(transactOps) > 0 {
		if transactOps[0] == nil {
			return nil, fmt.Errorf("transactOpts cannot be nil")
		}
		erc721Interactions.TxOptsFn = func() (*bind.TransactOpts, error) {
			return transactOps[0], nil
		}
	}

	return erc721Interactions, nil
}

// GetAddress returns the NFT contract address.
func (d *ERC721Interactions) GetAddress() common.Address {
	return d.nftAddress
}

// GetSession returns the current session used for NFT interactions.
func (d *ERC721Interactions) GetSession() transaction.Session {
	return d.session
}

// GetBalance retrieves the balance of NFTs for the associated address.
func (d *ERC721Interactions) GetBalance() (*big.Int, error) {
	balance, err := transaction.Call(
		d.session,
		d.session.erc271.PackBalanceOf(d.Address),
		d.session.erc271.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("nft.BalanceOf()", err)
	}
	return balance, nil
}

// TransferTo transfers a specific token to another address after verifying ownership.
func (d *ERC721Interactions) TransferTo(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.session.erc271.PackSafeTransferFrom(d.Address, to, tokenID),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("nft.TransferFrom()", err)
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
			if strings.Contains(err.Error(), hex.ErrZeroAddress.Error()) {
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
	supply, err := transaction.Call(
		d.session,
		d.session.erc271.PackTotalSupply(),
		d.session.erc271.UnpackTotalSupply,
	)
	if err != nil {
		return nil, d.callError("nft.TotalSupply()", err)
	}
	return supply, nil
}

// BalanceOf retrieves the NFT balance for a given owner.
func (d *ERC721Interactions) BalanceOf(owner common.Address) (*big.Int, error) {
	balance, err := transaction.Call(
		d.session,
		d.session.erc271.PackBalanceOf(owner),
		d.session.erc271.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("nft.BalanceOf()", err)
	}
	return balance, nil
}

// OwnerOf retrieves the owner of a specific token.
func (d *ERC721Interactions) OwnerOf(tokenID *big.Int) (common.Address, error) {
	owner, err := transaction.Call(
		d.session,
		d.session.erc271.PackOwnerOf(tokenID),
		d.session.erc271.UnpackOwnerOf,
	)
	if err != nil {
		return common.Address{}, d.callError("nft.OwnerOf()", err)
	}
	return owner, nil
}

// Approve approves an address to transfer a specific token.
func (d *ERC721Interactions) Approve(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.session.erc271.PackApprove(to, tokenID),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("nft.Approve()", err)
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
	name, err := transaction.Call(
		d.session,
		d.session.erc271.PackName(),
		d.session.erc271.UnpackName,
	)

	if err != nil {
		return "", d.callError("nft.Name()", err)
	}
	return name, nil
}

// Symbol returns the symbol of the NFT.
func (d *ERC721Interactions) Symbol() (string, error) {
	symbol, err := transaction.Call(
		d.session,
		d.session.erc271.PackSymbol(),
		d.session.erc271.UnpackSymbol,
	)

	if err != nil {
		return "", d.callError("nft.Symbol()", err)
	}
	return symbol, nil
}

// TokenURI returns the URI of the NFT.
func (d *ERC721Interactions) TokenURI(tokenID *big.Int) (string, error) {
	uri, err := transaction.Call(
		d.session,
		d.session.erc271.PackTokenURI(tokenID),
		d.session.erc271.UnpackTokenURI,
	)

	if err != nil {
		return "", d.callError("nft.TokenURI()", err)
	}
	return uri, nil
}

// GetApproved returns the approved address for a specific token.
func (d *ERC721Interactions) GetApproved(tokenID *big.Int) (common.Address, error) {
	approved, err := transaction.Call(
		d.session,
		d.session.erc271.PackGetApproved(tokenID),
		d.session.erc271.UnpackGetApproved,
	)
	if err != nil {
		return common.Address{}, d.callError("nft.GetApproved()", err)
	}
	return approved, nil
}
