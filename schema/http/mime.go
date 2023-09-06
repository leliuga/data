package http

import (
	"bytes"
	"errors"
	"strings"
)

// Common HTTP mimes
const (
	MimeInvalid Mime = iota //
	MimeTextXML
	MimeTextHTML
	MimeTextPlain
	MimeTextJavaScript
	MimeApplicationXML
	MimeApplicationJSON

	MimeApplicationForm
	MimeOctetStream
	MimeMultipartForm

	MimeTextXMLCharsetUTF8
	MimeTextHTMLCharsetUTF8
	MimeTextPlainCharsetUTF8
	MimeTextJavaScriptCharsetUTF8
	MimeApplicationXMLCharsetUTF8
	MimeApplicationJSONCharsetUTF8
)

var (
	// MimeNames is a map of Mime to string.
	MimeNames = map[Mime]string{
		MimeTextXML:         "text/xml",
		MimeTextHTML:        "text/html",
		MimeTextPlain:       "text/plain",
		MimeTextJavaScript:  "text/javascript",
		MimeApplicationXML:  "application/xml",
		MimeApplicationJSON: "application/json",

		MimeApplicationForm: "application/x-www-form-urlencoded",
		MimeOctetStream:     "application/octet-stream",
		MimeMultipartForm:   "multipart/form-data",

		MimeTextXMLCharsetUTF8:         "text/xml; charset=utf-8",
		MimeTextHTMLCharsetUTF8:        "text/html; charset=utf-8",
		MimeTextPlainCharsetUTF8:       "text/plain; charset=utf-8",
		MimeTextJavaScriptCharsetUTF8:  "text/javascript; charset=utf-8",
		MimeApplicationXMLCharsetUTF8:  "application/xml; charset=utf-8",
		MimeApplicationJSONCharsetUTF8: "application/json; charset=utf-8",
	}

	// ErrMimeInvalid is returned if the HTTP mime is invalid.
	ErrMimeInvalid = errors.New("invalid http mime")
)

// String mime to string
func (m Mime) String() string {
	return MimeNames[m]
}

// MarshalJSON mime to json
func (m Mime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

// UnmarshalJSON mime from json
func (m *Mime) UnmarshalJSON(b []byte) error {
	*m = ParseMime(string(bytes.Trim(b, `"`)))

	return nil
}

// ParseMime parses mime string.
func ParseMime(name string) Mime {
	name = strings.ToUpper(name)
	for k, v := range MimeNames {
		if v == name {
			return k
		}
	}

	return MimeInvalid
}

// MustParseMime parses mime string or panics.
func MustParseMime(name string) Mime {
	v := ParseMime(name)
	if v == MimeInvalid {
		panic(ErrMimeInvalid)
	}

	return v
}
