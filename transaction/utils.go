package transaction

// DefaultUnpacker is a default implementation that returns zero values.
func DefaultUnpacker([]byte) (byte, error) { return 0, nil }
