package data

type (
	// Kind defines a data kind.
	Kind uint8

	// IModel defines a model interface.
	IModel interface {
		// Validate makes `Model` validatable by implementing [validation.Validatable] interface.
		Validate() error
	}
)
