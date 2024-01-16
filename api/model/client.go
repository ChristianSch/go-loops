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

// interface for the API client
type ApiClient interface {
	ContactAPI() ContactAPI
	PrepareRequest(payload io.Reader, method HttpMethod, subUrl string) (*http.Request, error)
}

type ContactAPI interface {
	CreateContact(CreateContactData) (*CreateContactSuccess, error)
}
