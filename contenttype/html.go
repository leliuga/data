package contenttype

import (
	"bytes"
	"fmt"
	"io"
)

// Marshal returns a reader for the given value.
func (ht *HtmlType) Marshal(value any) (io.Reader, error) {
	return bytes.NewBufferString(fmt.Sprintf("%v", value)), nil
}

// Unmarshal parses the given reader and stores the result in the value pointed to by value.
func (ht *HtmlType) Unmarshal(r io.Reader, value any) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	value = string(b)

	return nil
}
