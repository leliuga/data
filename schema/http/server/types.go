package server

import (
	"github.com/leliuga/data"
	"github.com/leliuga/data/schema/http"
	"github.com/leliuga/validation"
)

type (
	// Endpoint represents the http server endpoint.
	Endpoint struct {
		validation.Validatable `json:"-"`
		Name                   string           `json:"name"          yaml:"Name"`
		Method                 http.Method      `json:"method"        yaml:"Method"`
		Path                   string           `json:"path"          yaml:"Path"`
		Description            string           `json:"description"   yaml:"Description"`
		Documentation          string           `json:"documentation" yaml:"Documentation"`
		Deprecated             string           `json:"deprecated"    yaml:"Deprecated"`
		Labels                 data.Map[string] `json:"labels"        yaml:"Labels"`
		IsPublic               bool             `json:"is_public"     yaml:"IsPublic"`
		IsStatic               bool             `json:"is_static"     yaml:"IsStatic"`
		Request                data.IModel      `json:"request"       yaml:"Request"`
		Response               data.IModel      `json:"response"      yaml:"Response"`
		Handlers               []Handler        `json:"-"`
	}

	// Endpoints represents the http server endpoints.
	Endpoints []*Endpoint

	// Handler represents the http server endpoint handler.
	Handler func(IContext) error

	// IContext represents the http server request context.
	IContext interface {
		// Validate makes Endpoint validatable by implementing [validation.Validatable] interface.
		Validate() error
	}
)
