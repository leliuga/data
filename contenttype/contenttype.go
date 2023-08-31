package contenttype

import (
	"bytes"
	"errors"
	"strings"
)

const (
	Invalid ContentType = iota
	Form
	Html
	Json
	MsgPack
	Text
	Yaml
)

var (
	// Names is a map of content type names to content type values.
	Names = map[ContentType]string{
		Form:    "application/x-www-form-urlencoded",
		Html:    "text/html",
		Json:    "application/json",
		MsgPack: "application/msgpack",
		Text:    "text/plain",
		Yaml:    "application/yaml",
	}

	// Set is a map of content type values to content type marshal and unmarshal.
	Set = map[ContentType]IContentType{
		Form:    &FormType{},
		Html:    &HtmlType{},
		Json:    &JsonType{},
		MsgPack: &MsgPackType{},
		Text:    &TextType{},
		Yaml:    &YamlType{},
	}

	// ErrInvalid is returned when the content type is invalid.
	ErrInvalid = errors.New("invalid content type")
)

// String content type to string
func (ct ContentType) String() string {
	return Names[ct]
}

// MarshalJSON content type to json
func (ct ContentType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.String() + `"`), nil
}

// UnmarshalJSON content type from json
func (ct *ContentType) UnmarshalJSON(b []byte) error {
	*ct = Parse(string(bytes.Trim(b, `"`)))

	return nil
}

// Parse parses content type string.
func Parse(name string) ContentType {
	for k, v := range Names {
		if strings.HasPrefix(name, v) {
			return k
		}
	}

	return Invalid
}

// MustParse parses content type string or panics.
func MustParse(name string) ContentType {
	v := Parse(name)
	if v == Invalid {
		panic(ErrInvalid)
	}

	return v
}
