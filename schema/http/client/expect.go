package client

import (
	"net/http"

	"github.com/pkg/errors"
)

// NewExpect creates a new Expect.
func NewExpect(headers map[string]string) *Expect {
	return &Expect{
		Status:  http.StatusOK,
		Headers: headers,
	}
}

// Validate validates the given response.
func (e *Expect) Validate(response *http.Response) error {
	if err := e.ValidateStatus(response); err != nil {
		return err
	}

	if err := e.ValidateHeader(response); err != nil {
		return err
	}

	return nil
}

// ValidateStatus validates the status code of the given response.
func (e *Expect) ValidateStatus(response *http.Response) error {
	if response.StatusCode != e.Status {
		return errors.Errorf("unexpected status code: %d (expected: %d)", response.StatusCode, e.Status)
	}

	return nil
}

// ValidateHeader validates the headers of the given response.
func (e *Expect) ValidateHeader(response *http.Response) error {
	for k, v := range e.Headers {
		if value := response.Header.Get(k); value != v {
			return errors.Errorf("unexpected header: %s=%s (expected: %s)", k, value, v)
		}
	}

	return nil
}
