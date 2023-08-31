package contenttype

import (
	"bytes"
	"io"

	"github.com/vmihailenco/msgpack/v5"
)

// Marshal returns a reader for the given value.
func (mt *MsgPackType) Marshal(value any) (io.Reader, error) {
	buffer := bytes.NewBuffer(nil)
	if err := msgpack.NewEncoder(buffer).Encode(value); err != nil {
		return nil, err
	}

	return buffer, nil
}

// Unmarshal parses the given reader and stores the result in the value pointed to by value.
func (mt *MsgPackType) Unmarshal(r io.Reader, value any) error {
	return msgpack.NewDecoder(r).Decode(value)
}
