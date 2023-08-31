package database

import (
	"fmt"
	"strings"

	"github.com/leliuga/data"
	"github.com/leliuga/data/constants"
	"github.com/leliuga/validation"
	"golang.org/x/exp/maps"
)

func NewColumn(kind data.Kind, name, nativeType string) *Column {
	return &Column{
		Kind:       kind,
		Name:       name,
		NativeKind: nativeType,
		Charset:    constants.DefaultCharset,
		Creatable:  true,
		Updatable:  true,
		Readable:   true,
	}
}

// Validate makes `Column` validatable by implementing [validation.Validatable] interface.
func (c *Column) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Kind, validation.In(validation.ToAnySliceFromMapKeys(data.KindNames)...).Error(fmt.Sprintf("A kind value must be one of: %s", strings.Join(maps.Values(data.KindNames), ", ")))),
		validation.Field(&c.Name, validation.Required, validation.Length(1, 63), validation.Match(constants.NameRegex).Error(constants.InvalidName)),
		validation.Field(&c.NativeKind, validation.Required),
	)
}
