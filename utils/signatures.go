// Package utils provides common utilities and constants for Ethereum contract interactions.
package utils

type Signature interface {
	GetHex() string
}
