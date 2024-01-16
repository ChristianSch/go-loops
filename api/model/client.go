package model

import (
	"io"
	"net/http"
)

// enum for http methods
type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

// ApiError is returned when the Loops API returns an error
type ApiError struct {
	Message string
}

const (
	// ApiErrorInvalidApiKey is returned when the API key is invalid
	ApiErrorInvalidApiKey = "invalid api key"
)

// interface for the API client
type ApiClient interface {
	// ContactAPI returns the ContactAPI
	ContactAPI() ContactAPI
	// PrepareRequest prepares a http request for the Loops API
	PrepareRequest(payload io.Reader, method HttpMethod, subUrl string) (*http.Request, error)
}

// interface for the contact API
type ContactAPI interface {
	// CreateContact creates a new contact
	CreateContact(CreateContactData) (*CreateContactSuccess, error)
}
