package http

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
)

// Common HTTP methods, these are defined in RFC 7231 section 4.3.
const (
	MethodInvalid Method = iota //
	MethodGet                   // RFC 7231, 4.3.1
	MethodPost                  // RFC 7231, 4.3.3
	MethodPut                   // RFC 7231, 4.3.4
	MethodPatch                 // RFC 5789
	MethodDelete                // RFC 7231, 4.3.5
	MethodHead                  // RFC 7231, 4.3.2
	MethodOptions               // RFC 7231, 4.3.7
	MethodTrace                 // RFC 7231, 4.3.8
)

var (
	// MethodNames is a map of Method to string.
	MethodNames = map[Method]string{
		MethodGet:     http.MethodGet,
		MethodPost:    http.MethodPost,
		MethodPut:     http.MethodPut,
		MethodPatch:   http.MethodPatch,
		MethodDelete:  http.MethodDelete,
		MethodHead:    http.MethodHead,
		MethodOptions: http.MethodOptions,
		MethodTrace:   http.MethodTrace,
	}

	// ErrMethodInvalid is returned if the HTTP method is invalid.
	ErrMethodInvalid = errors.New("invalid method")
)

// String method to string
func (m Method) String() string {
	return MethodNames[m]
}

// MarshalJSON method to json
func (m Method) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

// UnmarshalJSON method from json
func (m *Method) UnmarshalJSON(b []byte) error {
	*m = ParseMethod(string(bytes.Trim(b, `"`)))

	return nil
}

// ParseMethod parses method string.
func ParseMethod(name string) Method {
	for k, v := range MethodNames {
		if v == strings.ToUpper(name) {
			return k
		}
	}

	return MethodInvalid
}

// MustParseMethod parses method string or panics.
func MustParseMethod(name string) Method {
	v := ParseMethod(name)
	if v == MethodInvalid {
		panic(ErrMethodInvalid)
	}

	return v
}
