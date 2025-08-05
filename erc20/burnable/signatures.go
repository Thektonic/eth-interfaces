// Package burnable provides functions to interact with ERC20 burnable properties.
package burnable

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

type ERC20BurnableSignatures string

const (
	Burn     ERC20BurnableSignatures = "burn(uint256)"
	BurnFrom ERC20BurnableSignatures = "burnFrom(address,uint256)"
)

func (s ERC20BurnableSignatures) GetHex() string {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(s)) // hash.Write never returns an error
	selector := hash.Sum(nil)[:4]
	return hex.EncodeToString(selector)
}
