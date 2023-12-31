package http

import (
	"bytes"
	"errors"
	"strings"
)

// Common HTTP methods, these are defined in RFC 7231 section 4.3.
const (
	MethodInvalid Method = iota //
	MethodGet                   // RFC 7231, 4.3.1
	MethodHead                  // RFC 7231, 4.3.2
	MethodPost                  // RFC 7231, 4.3.3
	MethodPut                   // RFC 7231, 4.3.4
	MethodDelete                // RFC 7231, 4.3.5
	MethodConnect               // RFC 7231, 4.3.6
	MethodOptions               // RFC 7231, 4.3.7
	MethodTrace                 // RFC 7231, 4.3.8
	MethodPatch                 // RFC 5789
)

var (
	// MethodNames is a map of Method to string.
	MethodNames = map[Method]string{
		MethodGet:     "GET",
		MethodHead:    "HEAD",
		MethodPost:    "POST",
		MethodPut:     "PUT",
		MethodDelete:  "DELETE",
		MethodConnect: "CONNECT",
		MethodOptions: "OPTIONS",
		MethodTrace:   "TRACE",
		MethodPatch:   "PATCH",
	}

	// ErrMethodInvalid is returned if the HTTP method is invalid.
	ErrMethodInvalid = errors.New("invalid http method")
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
	name = strings.ToUpper(name)
	for k, v := range MethodNames {
		if v == name {
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
