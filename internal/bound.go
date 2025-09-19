package internal

import "github.com/ghosind/collection"

// CheckIndex checks if the given index is in the range [0, size). If not, it panics with an "index
// out of bounds" error.
func CheckIndex(i, size int) {
	if i < 0 || i >= size {
		panic(collection.ErrOutOfBounds)
	}
}
