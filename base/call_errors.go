package base

import (
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

// GenCallError generates a call error handler that wraps contract call errors with additional context.
func GenCallError(
	kind string,
	buildError func(any) error,
	unpackError func(raw []byte) (any, error),
) func(field string, err error) error {
	return func(field string, err error) error {
		if err == nil {
			return nil
		}
		errBytes, success := ethclient.RevertErrorData(err)
		if success {
			data, unpackErr := unpackError(errBytes)
			if unpackErr != nil {
				return fmt.Errorf("failed to unpack error data: %w", unpackErr)
			}
			if customErr := buildError(data); customErr != nil {
				return WrapCallError(kind, field, customErr)
			}
		}
		return err
	}
}
