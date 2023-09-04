package client

import (
	"github.com/leliuga/data"
	"github.com/leliuga/data/schema/http"
	"github.com/leliuga/validation"
)

type (
	// Endpoint is an HTTP client endpoint.
	Endpoint struct {
		validation.Validatable `json:"-"`
		Name                   string           `json:"name"          yaml:"Name"`
		Method                 http.Method      `json:"method"        yaml:"Method"`
		Path                   string           `json:"path"          yaml:"Path"`
		Description            string           `json:"description"   yaml:"Description"`
		Documentation          string           `json:"documentation" yaml:"Documentation"`
		Deprecated             string           `json:"deprecated"    yaml:"Deprecated"`
		Labels                 data.Map[string] `json:"labels"        yaml:"Labels"`
		Headers                http.Headers     `json:"headers"       yaml:"Headers"`
		Payload                any              `json:"payload"       yaml:"Payload"`
		Expect                 *Expect          `json:"expect"        yaml:"Expect"`
	}

	// Expect is an HTTP response expectation.
	Expect struct {
		Status  http.Status  `json:"status"  yaml:"Status"`
		Headers http.Headers `json:"headers" yaml:"Headers"`
	}

	// Endpoints represents the http client endpoints.
	Endpoints []*Endpoint
)
