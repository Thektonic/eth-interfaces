package transaction

import "fmt"

// DefaultUnpacker is a default implementation that returns zero values.
func DefaultUnpacker([]byte) (byte, error) { return 0, nil }

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *JsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}
