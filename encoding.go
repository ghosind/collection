package collection

import (
	"encoding/json"
	"fmt"
)

// Stringer is an interface that wraps the basic String method.
type Stringer fmt.Stringer

// JSONMarshaler is an interface that wraps the basic MarshalJSON method.
type JSONMarshaler json.Marshaler

// JSONUnmarshaler is an interface that wraps the basic UnmarshalJSON method.
type JSONUnmarshaler json.Unmarshaler
