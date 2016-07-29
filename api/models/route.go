package models

import (
	"errors"
	"net/http"
)

var (
	ErrRoutesCreate   = errors.New("Could not create route")
	ErrRoutesUpdate   = errors.New("Could not update route")
	ErrRoutesRemoving = errors.New("Could not remove route from datastore")
	ErrRoutesGet      = errors.New("Could not get route from datastore")
	ErrRoutesList     = errors.New("Could not list routes from datastore")
	ErrRoutesNotFound = errors.New("Route not found")
)

type Routes []*Route

type Route struct {
	Name          string      `json:"name"`
	AppName       string      `json:"appname"`
	Path          string      `json:"path"`
	Image         string      `json:"image"`
	Type          string      `json:"type,omitempty"`
	ContainerPath string      `json:"container_path,omitempty"`
	Headers       http.Header `json:"headers,omitempty"`
}

var (
	ErrRoutesValidationName    = errors.New("Missing route Name")
	ErrRoutesValidationImage   = errors.New("Missing route Image")
	ErrRoutesValidationAppName = errors.New("Missing route AppName")
	ErrRoutesValidationPath    = errors.New("Missing route Path")
)

func (r *Route) Validate() error {
	if r.Name == "" {
		return ErrRoutesValidationName
	}

	if r.Image == "" {
		return ErrRoutesValidationImage
	}

	if r.AppName == "" {
		return ErrRoutesValidationAppName
	}

	if r.Path == "" {
		return ErrRoutesValidationPath
	}

	return nil
}

type RouteFilter struct {
	Path    string
	AppName string
}