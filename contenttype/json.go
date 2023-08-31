package contenttype

import (
	"bytes"
	"io"

	"github.com/goccy/go-json"
)

// Marshal returns a reader for the given value.
func (jt *JsonType) Marshal(value any) (io.Reader, error) {
	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(value); err != nil {
		return nil, err
	}

	return buffer, nil
}

// Unmarshal parses the given reader and stores the result in the value pointed to by value.
func (jt *JsonType) Unmarshal(r io.Reader, value any) error {
	return json.NewDecoder(r).Decode(value)
}
