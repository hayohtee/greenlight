package data

import (
	"fmt"
	"strconv"
)

// Runtime A custom type which has underlying type int32
type Runtime int64

// MarshalJSON satisfies the json.Marshaler interface and return the
// JSON-encoded value for the movie runtime ("<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
