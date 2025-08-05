// Package utils provides common utilities and constants for Ethereum contract interactions.
package utils

import (
	"errors"
	"math/big"
)

var (
	IERC721InterfaceID  = [4]byte{0x80, 0xac, 0x58, 0xcd}
	IERC20InterfaceID   = [4]byte{0x36, 0x37, 0x2b, 0x07}
	IERC1155InterfaceID = [4]byte{0xd9, 0xb6, 0x7a, 0x26}
)

var (
	MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))
)

var (
	// ErrZeroAddress is returned when the zero address is used
	ErrZeroAddress = errors.New("TransferToZeroAddress")
)
