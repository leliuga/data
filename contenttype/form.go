package contenttype

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"reflect"
)

// Marshal returns a reader for the given value.
func (ft *FormType) Marshal(value any) (io.Reader, error) {
	v := ""
	switch t := value.(type) {
	case string:
		v = t
	case url.Values:
		v = t.Encode()
	default:
		return nil, fmt.Errorf("a data type %s is invalid", reflect.TypeOf(t))
	}

	return bytes.NewBufferString(v), nil
}

// Unmarshal parses the given reader and stores the result in the value pointed to by value.
func (ft *FormType) Unmarshal(r io.Reader, value any) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	v, err := url.ParseQuery(string(b))
	if err != nil {
		return err
	}

	value = v

	return nil
}
