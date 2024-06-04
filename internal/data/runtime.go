package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Runtime A custom type which has underlying type int32
type Runtime int32

// ErrInvalidRuntimeFormat is an error that indicates there is a problem
// decoding the Runtime JSON value.
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// MarshalJSON satisfies the json.Marshaler interface and return the
// JSON-encoded value for the movie runtime ("<runtime> mins").
func (r *Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

// UnmarshalJSON satisfies the json.Unmarshaler interface and
// decode JSON value for the field Runtime or throw ErrInvalidRuntimeFormat
// if there is an error decoding the JSON value.
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}
