// Package royalties provides functions to interact with ERC721 royalty properties.
package royalties

import (
	"math/big"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/Thektonic/eth-interfaces/transaction"
)

// IERC721RoyaltiesInteractions wraps interactions with the IERC721Royalties contract.
type IERC721RoyaltiesInteractions struct {
	*nft.ERC721Interactions
	ierc721Royalties *inferences.Ierc721
	callError        func(string, error) *base.CallError
}

// NewERC721RoyaltiesInteractions creates a new instance of IERC721RoyaltiesInteractions.
func NewERC721RoyaltiesInteractions(
	baseIERC721 *nft.ERC721Interactions,
	signatures []IERC721RoyaltiesSignature,
) (*IERC721RoyaltiesInteractions, error) {
	var converted []hex.Signature
	for _, sig := range signatures {
		converted = append(converted, sig)
	}

	err := baseIERC721.CheckSignatures(baseIERC721.GetAddress(), converted)
	if err != nil {
		return nil, customerrors.WrapinterfacingError("ierc721Royalties", err)
	}

	ierc721Royalties := inferences.NewIerc721()

	callError := func(_ string, err error) *base.CallError {
		return baseIERC721.WrapCallError(inferences.Ierc721MetaData.ABI, "nft.RoyaltyInfo()", err)
	}

	return &IERC721RoyaltiesInteractions{baseIERC721, ierc721Royalties, callError}, nil
}

// RoyaltiesInfos retrieves the royalty information for a given token and sale price.
func (e *IERC721RoyaltiesInteractions) RoyaltiesInfos(tokenID *big.Int, salePrice *big.Int) (inferences.RoyaltyInfoOutput, error) {
	rInfos, err := transaction.Call(
		e.GetSession(),
		e.ierc721Royalties.PackRoyaltyInfo(tokenID, salePrice),
		e.ierc721Royalties.UnpackRoyaltyInfo,
	)

	if err != nil {
		return inferences.RoyaltyInfoOutput{}, e.callError("nft.RoyaltyInfo()", err)
	}
	return rInfos, nil
}
