package client

import (
	"fmt"
	"strings"

	"github.com/leliuga/data/constants"
	"github.com/leliuga/data/schema/http"
	"github.com/leliuga/validation"
	"golang.org/x/exp/maps"
)

// Validate makes Endpoint validatable by implementing [validation.Validatable] interface.
func (e *Endpoint) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Name, validation.Required, validation.Length(1, 63), validation.Match(constants.NameRegex).Error(constants.InvalidName)),
		validation.Field(&e.Method, validation.Required, validation.In(validation.ToAnySliceFromMapKeys(http.MethodNames)...).Error(fmt.Sprintf("A method value must be one of: %s", strings.Join(maps.Values(http.MethodNames), ", ")))),
		validation.Field(&e.Path, validation.Required),
	)
}
