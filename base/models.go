package base

import (
	"github.com/Thektonic/eth-interfaces/inferences/ERC20Burnable"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721A"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721Complete"
	"github.com/Thektonic/eth-interfaces/inferences/nftPositionManager"
	"github.com/ethereum/go-ethereum/common"
)

type Interactions[
	T *nftPositionManager.NftPositionManagerSession |
		*ERC721A.ERC721ASession |
		*ERC721Complete.ERC721CompleteSession |
		*ERC20Burnable.ERC20BurnableSession,
] struct {
	*BaseInteractions
	Session   T
	Address   common.Address
	CallError func(string, error) *CallError
}

// GetNFTAddress returns the NFT contract address.
func (d *Interactions[T]) GetAddress() common.Address {
	return d.Address
}

func (d *Interactions[T]) GetSession() T {
	return d.Session
}

func (d *Interactions[T]) GetBaseInteractions() *BaseInteractions {
	return d.BaseInteractions
}
