package contenttype

import (
	"io"
)

type (
	// FormType is a form type.
	FormType struct{}

	// HtmlType is a html type.
	HtmlType struct{}

	// JsonType is a json type.
	JsonType struct{}

	// MsgPackType is a msgpack type.
	MsgPackType struct{}

	// TextType is a text type.
	TextType struct{}

	// YamlType is a yaml type.
	YamlType struct{}

	// ContentType is a content type.
	ContentType uint8

	// IContentType is a content type interface.
	IContentType interface {
		// Marshal marshals the value to a reader.
		Marshal(any) (io.Reader, error)

		// Unmarshal unmarshals the reader to a value.
		Unmarshal(io.Reader, any) error
	}
)
