package data

import (
	"bytes"
	"errors"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Default values for the date/time formats
const (
	DefaultDateFormat     = "2006-01-02"
	DefaultDateTimeFormat = "2006-01-02 15:04:05"
	DefaultTimeFormat     = "15:04:05"
)

const (
	KindInvalid   Kind = iota + 1 //
	KindReference                 // The `reference`scalar kind represents a reference to other structure Columns
	KindBoolean                   // The `boolean`  scalar kind represents                                                                        value `true` or `false`
	KindDateTime                  // The `datetime` scalar kind represents a date and time pairing in UTC                                         value `2023-01-25 10:10:10`
	KindDate                      // The `date`     scalar kind represents a date in UTC                                                          value `2023-01-25`
	KindFloat32                   // The `float32`  scalar kind represents an 32-bit signed double-precision floating-point                       value '1.2345'
	KindFloat64                   // The `float64`  scalar kind represents an 64-bit signed double-precision floating-point                       value '1.2345'
	KindID                        // The `id`       scalar kind represents a 128-bit hexadecimal                                                  value '7f9c24e8-3b12-4fef-91e0-56a2d5a246ec'
	KindInet                      // The `inet`     scalar kind represents an IPv4 or IPv6 address                                                value '192.168.0.1'
	KindUInt8                     // The `uint8`    scalar kind represents an unsigned  8-bit integers (0 to 255)                                 value '12345'
	KindUInt16                    // The `uint16`   scalar kind represents an unsigned 16-bit integers (0 to 65535)                               value '12345'
	KindUInt32                    // The `uint32`   scalar kind represents an unsigned 32-bit integers (0 to 4294967295)                          value '12345'
	KindUInt64                    // The `uint64`   scalar kind represents an unsigned 64-bit integers (0 to 18446744073709551615)                value '12345'
	KindInt8                      // The `int8`     scalar kind represents a signed  8-bit integers (-128 to 127)                                 value '-12'
	KindInt16                     // The `int16`    scalar kind represents a signed 16-bit integers (-32768 to 32767)                             value '-12'
	KindInt32                     // The `int32`    scalar kind represents a signed 32-bit integers (-2147483648 to 2147483647)                   value '-12'
	KindInt64                     // The `int64`    scalar kind represents a signed 64-bit integers (-9223372036854775808 to 9223372036854775807) value '-12'
	KindString                    // The `string`   scalar kind represents textual data a UTF‐8 character sequence                                value 'a1b2c3'
	KindTime                      // The `time`     scalar kind represents a time in UTC                                                          value `01:23:45.123456`
)

var (
	KindNames = map[Kind]string{
		KindReference: "Reference",
		KindBoolean:   "Boolean",
		KindDateTime:  "Datetime",
		KindDate:      "Date",
		KindFloat32:   "Float32",
		KindFloat64:   "Float64",
		KindID:        "Id",
		KindInet:      "Inet",
		KindUInt8:     "Uint8",
		KindUInt16:    "Uint16",
		KindUInt32:    "Uint32",
		KindUInt64:    "Uint64",
		KindInt8:      "Int8",
		KindInt16:     "Int16",
		KindInt32:     "Int32",
		KindInt64:     "Int64",
		KindString:    "String",
		KindTime:      "Time",
	}

	// ErrKindInvalid is returned when the data kind is invalid.
	ErrKindInvalid = errors.New("invalid data kind")
)

// String data kind to string
func (k Kind) String() string {
	return KindNames[k]
}

// MarshalJSON data kind to json
func (k Kind) MarshalJSON() ([]byte, error) {
	return []byte(`"` + k.String() + `"`), nil
}

// UnmarshalJSON data kind from json
func (k *Kind) UnmarshalJSON(b []byte) error {
	*k = ParseKind(string(bytes.Trim(b, `"`)))

	return nil
}

// ParseKind parses data kind string.
func ParseKind(name string) Kind {
	for k, v := range KindNames {
		if v == strings.ToLower(name) {
			return k
		}
	}

	return KindInvalid
}

// MustParseKind parses data kind string or panics.
func MustParseKind(name string) Kind {
	v := ParseKind(name)
	if v == KindInvalid {
		panic(ErrKindInvalid)
	}

	return v
}

func DetectValueKind(value any, convert bool) Kind {
	switch value.(type) {
	case bool:
		{
			return KindBoolean
		}
	case float64:
		{
			v := value.(float64)
			if v <= math.MaxFloat32 && v >= -math.MaxFloat32 {
				return KindFloat32
			}

			return KindFloat64
		}
	case int:
		{
			v := value.(int)
			if v <= math.MaxInt8 && v >= math.MinInt8 {
				return KindInt8
			} else if v <= math.MaxInt16 && v >= math.MinInt16 {
				return KindInt16
			} else if v <= math.MaxInt32 && v >= math.MinInt32 {
				return KindInt32
			}
			return KindInt64
		}
	case string:
		{
			if convert {
				v := strings.ToLower(strings.TrimSpace(value.(string)))
				if _, err := strconv.ParseInt(v, 10, 8); err == nil {
					return KindInt8
				} else if _, err = strconv.ParseInt(v, 10, 16); err == nil {
					return KindInt16
				} else if _, err = strconv.ParseInt(v, 10, 32); err == nil {
					return KindInt32
				} else if _, err = strconv.ParseInt(v, 10, 64); err == nil {
					return KindInt64
				} else if _, err = strconv.ParseFloat(v, 32); err == nil {
					return KindFloat32
				} else if _, err = strconv.ParseFloat(v, 64); err == nil {
					return KindFloat64
				} else if _, err = strconv.ParseBool(v); err == nil {
					return KindBoolean
				} else if _, err = time.Parse(DefaultDateTimeFormat, v); err == nil {
					return KindDateTime
				} else if _, err = time.Parse(DefaultDateFormat, v); err == nil {
					return KindDate
				} else if _, err = time.Parse(DefaultTimeFormat, v); err == nil {
					return KindTime
				} else if _, err = uuid.Parse(v); err == nil {
					return KindID
				}

				_, _, ok := net.ParseCIDR(v)
				if ok == nil || net.ParseIP(v) != nil {
					return KindInet
				}
			}

			return KindString
		}
	}

	return KindInvalid
}
