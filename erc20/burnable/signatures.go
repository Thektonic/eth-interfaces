// Package burnable provides functions to interact with ERC20 burnable properties.
package burnable

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// ERC20BurnableSignatures represents function signatures for ERC20 burnable token operations
type ERC20BurnableSignatures string

const (
	// Burn represents the burn function signature for burning tokens
	Burn ERC20BurnableSignatures = "burn(uint256)"
	// BurnFrom represents the burnFrom function signature for burning tokens from another address
	BurnFrom ERC20BurnableSignatures = "burnFrom(address,uint256)"
)

// computeHash returns the Keccak256 hash of the function signature
func (s ERC20BurnableSignatures) computeHash() []byte {
	hash := crypto.NewKeccakState()
	_, _ = hash.Write([]byte(s)) // hash.Write never returns an error
	return hash.Sum(nil)
}

// GetHex returns the hex representation of the function signature
func (s ERC20BurnableSignatures) GetHex() string {
	return hex.EncodeToString(s.computeHash())
}

func (s ERC20BurnableSignatures) String() string {
	return string(s)
}

// GetSelector returns the Keccak256 hash selector for the ERC20 burnable signature
func (s ERC20BurnableSignatures) GetSelector() []byte {
	return s.computeHash()[:4]
}
