package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leliuga/data"
	"github.com/leliuga/data/labels"
	"github.com/leliuga/data/schema/http"
	"github.com/leliuga/validation"
)

type (
	// Endpoint represents the http server endpoint.
	Endpoint struct {
		validation.Validatable `json:"-"`
		Name                   string          `json:"name"          yaml:"Name"`
		Method                 http.Method     `json:"method"        yaml:"Method"`
		Path                   string          `json:"path"          yaml:"Path"`
		Description            string          `json:"description"   yaml:"Description"`
		Documentation          string          `json:"documentation" yaml:"Documentation"`
		Deprecated             string          `json:"deprecated"    yaml:"Deprecated"`
		Labels                 labels.Labels   `json:"labels"        yaml:"Labels"`
		IsPublic               bool            `json:"is_public"     yaml:"IsPublic"`
		IsStatic               bool            `json:"is_static"     yaml:"IsStatic"`
		Request                data.IModel     `json:"request"       yaml:"Request"`
		Response               data.IModel     `json:"response"      yaml:"Response"`
		Handlers               []fiber.Handler `json:"-"`
	}

	// Endpoints represents the http server endpoints.
	Endpoints []*Endpoint
)
