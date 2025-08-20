package transaction

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

type Session interface {
	CallOpts() *bind.CallOpts
	Instance() *bind.BoundContract
}

func Call[T any](s Session, calldata []byte, unpack func([]byte) (T, error)) (T, error) {
	return bind2.Call(s.Instance(), s.CallOpts(), calldata, unpack)
}
