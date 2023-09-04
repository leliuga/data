package http

type (
	// Header defines an HTTP header.
	Header uint16

	// Method defines an HTTP method.
	Method uint8

	// Status defines an HTTP status.
	Status uint16

	// Headers represents a collection of HTTP headers.
	Headers map[Header]string
)
