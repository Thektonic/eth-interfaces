// Package hex provides common utilities and constants for Ethereum contract interactions.
package hex

import (
	"errors"
	"math/big"
)

var (
	// IERC721InterfaceID is the interface ID for ERC721 tokens
	IERC721InterfaceID = [4]byte{0x80, 0xac, 0x58, 0xcd}
	// IERC20InterfaceID is the interface ID for ERC20 tokens
	IERC20InterfaceID = [4]byte{0x36, 0x37, 0x2b, 0x07}
	// IERC1155InterfaceID is the interface ID for ERC1155 tokens
	IERC1155InterfaceID = [4]byte{0xd9, 0xb6, 0x7a, 0x26}
)

const (
	// ErrorMethodIDLength is the length of Ethereum error method ID in bytes
	ErrorMethodIDLength = 4
	// DecimalBase is the base for decimal number parsing
	DecimalBase = 10
	// TestChainID is the chain ID used for testing
	TestChainID = 1337
	// TestGasLimit is the gas limit used for testing
	TestGasLimit = 9_000_000
	// Uint256BitSize is the bit size for uint256 values
	Uint256BitSize = 256
	// MaxUint256Offset is the offset used in MaxUint256 calculation
	MaxUint256Offset = 9
)

var (
	// MaxUint256 represents the maximum value for a uint256
	MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), Uint256BitSize), big.NewInt(MaxUint256Offset))
)

var (
	// ErrZeroAddress is returned when the zero address is used
	ErrZeroAddress = errors.New("TransferToZeroAddress")
)

// ParseEther converts a big.Int wei value to a float64 ether value
func ParseEther(a *big.Int) float64 {
	fa, _ := a.Float64()
	return fa / (1e18)
}
