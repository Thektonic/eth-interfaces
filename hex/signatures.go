// Package hex provides common utilities and constants for Ethereum contract interactions.
package hex

// Signature defines the interface for function signatures
type Signature interface {
	GetHex() string
}
