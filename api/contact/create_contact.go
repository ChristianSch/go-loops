package contact

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ChristianSch/go-loops/api/model"
)

const (
	subUrl = "/contacts/create"
)

// internal API responses
type createContractResponse struct {
	Id      *string `json:"id"`
	Success *bool   `json:"success"`
	Message *string `json:"message"`
}

// CreateContact creates a new contact
func (c ContactAPI) CreateContact(contact model.CreateContactData) (*model.CreateContactSuccess, error) {
	// first step is to marshal the input data
	// as we have the custom fields it's a bit more complicated
	cJson, err := contact.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := c.client.PrepareRequest(bytes.NewReader(cJson), "POST", subUrl)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resData createContractResponse

	err = json.NewDecoder(res.Body).Decode(&resData)
	if err != nil {
		return nil, err
	}

	if resData.Success != nil && !*resData.Success {
		if resData.Message != nil {
			return nil, errors.New(*resData.Message)
		} else {
			return nil, errors.New("unknown error")
		}
	}

	if resData.Id == nil {
		return nil, errors.New("missing contact id")
	}

	return &model.CreateContactSuccess{ContactId: *resData.Id}, nil
}
