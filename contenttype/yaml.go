package contenttype

import (
	"bytes"
	"io"

	"github.com/goccy/go-yaml"
)

// Marshal returns a reader for the given value.
func (yt *YamlType) Marshal(value any) (io.Reader, error) {
	buffer := bytes.NewBuffer(nil)
	if err := yaml.NewEncoder(buffer).Encode(value); err != nil {
		return nil, err
	}

	return buffer, nil
}

// Unmarshal parses the given reader and stores the result in the value pointed to by value.
func (yt *YamlType) Unmarshal(r io.Reader, value any) error {
	return yaml.NewDecoder(r).Decode(value)
}
