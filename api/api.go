package api

import (
	"io"
	"net/http"

	"github.com/ChristianSch/go-loops/api/contact"
	"github.com/ChristianSch/go-loops/api/model"
)

type LoopsApiClient struct {
	Url        string
	Token      string
	contactApi model.ContactAPI
}

// NewApiClient creates a new LoopsApiClient
// url: the url of the loops instance (e.g. https://app.loops.so/api/v1)
// token: the api token of the loops API (see https://app.loops.so/settings?page=api)
func NewApiClient(url string, token string) *LoopsApiClient {
	// remove trailing slash  of url
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	c := &LoopsApiClient{Url: url, Token: token}

	c.contactApi = contact.NewContactAPI(c)

	return c
}

func (c LoopsApiClient) getUrl(subUrl string) string {
	return c.Url + subUrl
}

// PrepareRequest prepares a http request for the Loops API
func (c LoopsApiClient) PrepareRequest(payload io.Reader, method model.HttpMethod, subUrl string) (*http.Request, error) {
	req, err := http.NewRequest(string(method), c.getUrl(subUrl), payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// ContactAPI returns the ContactAPI
func (c LoopsApiClient) ContactAPI() model.ContactAPI {
	return c.contactApi
}
