package collection

import "errors"

var (
	// ErrOutOfBounds indicates that the index is out of the valid range.
	ErrOutOfBounds = errors.New("index out of bounds")
)
