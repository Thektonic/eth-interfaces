package base

import (
	"fmt"
)

// CallError represents an error that occurred during a contract call
type CallError struct {
	Field string
	Err   error
}

func (e *CallError) Error() string {
	return fmt.Sprintf("call error on %s: %s", e.Field, e.Err.Error())
}

// WrapCallError wraps an error with contract call context information
func (i *Interactions) WrapCallError(abiString, field string, err error) *CallError {
	if err == nil {
		return nil
	}
	if customErr := i.ManageCustomContractError(abiString, err); customErr != nil {
		return &CallError{
			Field: field,
			Err:   customErr,
		}
	}
	return &CallError{
		Field: field,
		Err:   err,
	}
}

// Unwrap returns the underlying error
func (e *CallError) Unwrap() error { return e.Err }
