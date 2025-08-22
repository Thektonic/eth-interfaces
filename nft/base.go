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
	erc721   *inferences.Ierc721
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
	callError  func(string, error) error
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
		return nil, customerrors.WrapInterfacingError("CheckSignatures", err)
	}

	erc721 := inferences.NewIerc721()

	erc721ASession := session{
		erc721:   erc721,
		callOpts: &bind.CallOpts{Pending: true, From: baseInteractions.Address},
		instance: erc721.Instance(baseInteractions.Client, address),
	}

	callError := base.GenCallError("erc721", ParseError, erc721.UnpackError)

	erc721Interactions := &ERC721Interactions{
		baseInteractions,
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
		d.erc721.PackBalanceOf(d.Address),
		d.erc721.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("BalanceOf()", err)
	}
	return balance, nil
}

// TransferTo transfers a specific token to another address after verifying ownership.
func (d *ERC721Interactions) TransferTo(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.erc721.PackSafeTransferFrom(d.Address, to, tokenID),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("TransferFrom()", err)
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
		d.erc721.PackTotalSupply(),
		d.erc721.UnpackTotalSupply,
	)
	if err != nil {
		return nil, d.callError("TotalSupply()", err)
	}
	return supply, nil
}

// BalanceOf retrieves the NFT balance for a given owner.
func (d *ERC721Interactions) BalanceOf(owner common.Address) (*big.Int, error) {
	balance, err := transaction.Call(
		d.session,
		d.erc721.PackBalanceOf(owner),
		d.erc721.UnpackBalanceOf,
	)
	if err != nil {
		return nil, d.callError("BalanceOf()", err)
	}
	return balance, nil
}

// OwnerOf retrieves the owner of a specific token.
func (d *ERC721Interactions) OwnerOf(tokenID *big.Int) (common.Address, error) {
	owner, err := transaction.Call(
		d.session,
		d.erc721.PackOwnerOf(tokenID),
		d.erc721.UnpackOwnerOf,
	)
	if err != nil {
		return common.Address{}, d.callError("OwnerOf()", err)
	}
	return owner, nil
}

// Approve approves an address to transfer a specific token.
func (d *ERC721Interactions) Approve(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	tx, err := transaction.Transact(
		d,
		d.session,
		d.erc721.PackApprove(to, tokenID),
		transaction.DefaultUnpacker,
	)
	if err != nil {
		return nil, d.callError("Approve()", err)
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
		d.erc721.PackName(),
		d.erc721.UnpackName,
	)

	if err != nil {
		return "", d.callError("Name()", err)
	}
	return name, nil
}

// Symbol returns the symbol of the NFT.
func (d *ERC721Interactions) Symbol() (string, error) {
	symbol, err := transaction.Call(
		d.session,
		d.erc721.PackSymbol(),
		d.erc721.UnpackSymbol,
	)

	if err != nil {
		return "", d.callError("Symbol()", err)
	}
	return symbol, nil
}

// TokenURI returns the URI of the NFT.
func (d *ERC721Interactions) TokenURI(tokenID *big.Int) (string, error) {
	uri, err := transaction.Call(
		d.session,
		d.erc721.PackTokenURI(tokenID),
		d.erc721.UnpackTokenURI,
	)

	if err != nil {
		return "", d.callError("TokenURI()", err)
	}
	return uri, nil
}

// GetApproved returns the approved address for a specific token.
func (d *ERC721Interactions) GetApproved(tokenID *big.Int) (common.Address, error) {
	approved, err := transaction.Call(
		d.session,
		d.erc721.PackGetApproved(tokenID),
		d.erc721.UnpackGetApproved,
	)
	if err != nil {
		return common.Address{}, d.callError("GetApproved()", err)
	}
	return approved, nil
}

// ParseError parses raw contract errors into human-readable error messages for NFT/ERC721 operations.
func ParseError(rawErr any) error {
	switch e := rawErr.(type) {
	case *inferences.Ierc721ERC721OutOfBoundsIndex:
		return fmt.Errorf("ERC721OutOfBoundsIndex: %s, %s", e.Index.String(), e.Owner.Hex())
	case *inferences.Ierc721ERC721IncorrectOwner:
		return fmt.Errorf("ERC721IncorrectOwner: owner %s, spender %s, %s", e.Owner.Hex(), e.Sender.Hex(), e.TokenId.String())
	case *inferences.Ierc721ERC721InsufficientApproval:
		return fmt.Errorf("ERC721InsufficientApproval: %s, %s", e.Operator.String(), e.TokenId.String())
	case *inferences.Ierc721ERC721InvalidApprover:
		return fmt.Errorf("ERC721InvalidApprover: %s", e.Approver.Hex())
	case *inferences.Ierc721ERC721InvalidOperator:
		return fmt.Errorf("ERC721InvalidOperator: %s", e.Operator.String())
	case *inferences.Ierc721ERC721InvalidReceiver:
		return fmt.Errorf("ERC721InvalidReceiver:  %s", e.Receiver.Hex())
	case *inferences.Ierc721ERC721InvalidSender:
		return fmt.Errorf("ERC721InvalidSender: %s", e.Sender.String())
	case *inferences.Ierc721ERC721NonexistentToken:
		return fmt.Errorf("ERC721NonexistentToken: %s", e.TokenId.String())
	case *inferences.Ierc721OwnerQueryForNonexistentToken:
		return errors.New("OwnerQueryForNonexistentToken")
	case *inferences.Ierc721ApprovalCallerNotOwnerNorApproved:
		return errors.New("ApprovalCallerNotOwnerNorApproved")
	case *inferences.Ierc721ApprovalQueryForNonexistentToken:
		return errors.New("ApprovalQueryForNonexistentToken")
	case *inferences.Ierc721BalanceQueryForZeroAddress:
		return errors.New("BalanceQueryForZeroAddress")
	case *inferences.Ierc721MintToZeroAddress:
		return errors.New("MintToZeroAddress")
	case *inferences.Ierc721MintZeroQuantity:
		return errors.New("MintZeroQuantity")
	case *inferences.Ierc721TransferCallerNotOwnerNorApproved:
		return errors.New("TransferCallerNotOwnerNorApproved")
	case *inferences.Ierc721TransferFromIncorrectOwner:
		return errors.New("TransferFromIncorrectOwner")
	case *inferences.Ierc721TransferToNonERC721ReceiverImplementer:
		return errors.New("TransferToNonERC721ReceiverImplementer")
	case *inferences.Ierc721TransferToZeroAddress:
		return errors.New("TransferToZeroAddress")
	case *inferences.Ierc721URIQueryForNonexistentToken:
		return errors.New("URIQueryForNonexistentToken")
	case *inferences.Ierc721MintERC2309QuantityExceedsLimit:
		return errors.New("MintERC2309QuantityExceedsLimit")
	case *inferences.Ierc721OwnershipNotInitializedForExtraData:
		return errors.New("OwnershipNotInitializedForExtraData")
	default:
		return nil
	}
}
