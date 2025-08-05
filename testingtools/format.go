// Package testingtools provides common utilities and constants for Ethereum contract interactions.
package testingtools

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Thektonic/eth-interfaces/hex"
)

// FloatTo18z converts a float64 to an 18-decimal big.Int representation
func FloatTo18z(amount float64) *big.Int {
	s := fmt.Sprintf("%f", amount)
	a := strings.Split(s, ".")
	for i := 0; i < 18; i++ {
		if len(a) > 1 && len(a[1]) > i {
			a[0] += string(a[1][i])
		} else {
			a[0] += "0"
		}
	}
	b, ok := new(big.Int).SetString(a[0], hex.DecimalBase)
	if !ok {
		panic("Could not set big.Int string for value " + s)
	}
	return b
}
