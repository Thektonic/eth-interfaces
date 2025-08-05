// Package customerrors provides custom error types for contract interface operations.
package customerrors

import "fmt"

// InterfacingError represents an error that occurred during interface operations
type InterfacingError struct {
	InterfaceName string
	Err           error
}

func (e *InterfacingError) Error() string {
	return fmt.Sprintf("interface setup error function %s, error : %s", e.InterfaceName, e.Err.Error())
}

// WrapinterfacingError wraps an error with interface context information
func WrapinterfacingError(interfaceName string, err error) *InterfacingError {
	if err == nil {
		return nil
	}

	return &InterfacingError{
		InterfaceName: interfaceName,
		Err:           err,
	}
}

// UnWrap returns the underlying error
func (e *InterfacingError) UnWrap() error { return e.Err }
