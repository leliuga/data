package data

type (
	// Kind defines a data kind.
	Kind uint8

	// Map defines a map of key:value. It implements Map.
	Map[T any] map[string]T

	// IModel defines a model interface.
	IModel interface {
		// Validate makes `Model` validatable by implementing [validation.Validatable] interface.
		Validate() error
	}
)
