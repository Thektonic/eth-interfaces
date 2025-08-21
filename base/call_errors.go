package base

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// CallError represents an error that occurred during a contract call
type CallError struct {
	Kind   string // The kind of call, e.g., "call", "deploy", etc.
	Method string
	Err    error
}

func (e *CallError) Error() string {
	return fmt.Sprintf("%s.%s: %s", e.Kind, e.Method, e.Err.Error())
}

// WrapCallError wraps an error with contract call context information
func WrapCallError(kind, field string, err error) *CallError {
	if err == nil {
		return nil
	}

	return &CallError{
		Kind:   kind,
		Method: field,
		Err:    err,
	}
}

// Unwrap returns the underlying error
func (e *CallError) Unwrap() error { return e.Err }

func GenCallError(kind string, buildError func(any) error, UnpackError func(raw []byte) (any, error)) func(field string, err error) error {
	return func(field string, err error) error {
		errBytes, success := ethclient.RevertErrorData(err)
		if success {
			data, err := UnpackError(errBytes)
			if err != nil {
				return errors.New("failed to unpack error data")
			}
			return WrapCallError(kind, field, buildError(data))

		}
		return err
	}
}
