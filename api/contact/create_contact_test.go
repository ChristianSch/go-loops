package contact_test

import (
	"fmt"
	"testing"

	"github.com/ChristianSch/go-loops/api"
	"github.com/ChristianSch/go-loops/api/model"
	"github.com/jarcoal/httpmock"
)

func asPtr[T any](v T) *T {
	return &v
}

func TestCreateContactSucceeds(t *testing.T) {
	// mock the api client with a mock http client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// register catch-all responder for the loops api
	httpmock.RegisterMatcherResponder("POST", "http://loops.local/api/v1/contacts/create",
		httpmock.BodyContainsString(`"firstName":"John","lastName":"Doe","email":"john@doe.foo","source":"test","subscribed":true,"userGroup":"test"`),
		httpmock.NewStringResponder(200, `{"id": "123", "success": true}`))

	// create the contact api
	client := api.NewApiClient("http://loops.local/api/v1", "token")

	// Given
	// Valid input to create a user
	contact := model.CreateContactData{
		FirstName:  asPtr("John"),
		LastName:   asPtr("Doe"),
		Email:      "john@doe.foo",
		Source:     asPtr("test"),
		Subscribed: asPtr(true),
		UserGroup:  asPtr("test"),
	}

	// When
	// Creating a contact
	res, err := client.ContactAPI().CreateContact(contact)

	// Then
	// No error should be returned
	if err != nil {
		t.Error("Unexpected error:", err)
		return
	}

	fmt.Println("res:", res)

	// The contact id should be returned
	if res.ContactId != "123" {
		t.Error("Unexpected contact id:", res.ContactId)
	}
}

func TestCreateContactFails(t *testing.T) {
	// mock the api client with a mock http client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// register catch-all responder for the loops api
	httpmock.RegisterResponder("POST", "http://loops.local/api/v1/contacts/create",
		httpmock.NewStringResponder(401, `{"message": "something went wrong", "success": false}`))

	// create the contact api
	client := api.NewApiClient("http://loops.local/api/v1", "token")

	// Given
	// Valid input to create a user
	contact := model.CreateContactData{
		Email: "john@doe.foo",
	}

	// When
	// Creating a contact
	res, err := client.ContactAPI().CreateContact(contact)

	// Then
	// An error should be returned
	if res != nil || err == nil {
		t.Error("Unexpected error:", err)
	}

	// The message should be returned as the error message
	if err.Error() != "something went wrong" {
		t.Error("Unexpected error message:", err.Error())
	}
}

func TestCreateContactSucceedsWithCustomFields(t *testing.T) {
	// mock the api client with a mock http client
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// register catch-all responder for the loops api
	httpmock.RegisterMatcherResponder("POST", "http://loops.local/api/v1/contacts/create",
		httpmock.BodyContainsString(`"firstName":"John","lastName":"Doe","email":"john@doe.foo","source":"test","subscribed":true,"userGroup":"test","custom-1":"test","hasCustom":true`),
		httpmock.NewStringResponder(200, `{"id": "123", "success": true}`))

	// create the contact api
	client := api.NewApiClient("http://loops.local/api/v1", "token")

	// Given
	// Valid input to create a user
	contact := model.CreateContactData{
		FirstName:  asPtr("John"),
		LastName:   asPtr("Doe"),
		Email:      "john@doe.foo",
		Source:     asPtr("test"),
		Subscribed: asPtr(true),
		UserGroup:  asPtr("test"),
	}
	contact.CustomFields = map[string]interface{}{
		"custom-1":  "test",
		"hasCustom": true,
	}

	// When
	// Creating a contact
	res, err := client.ContactAPI().CreateContact(contact)

	// Then
	// No error should be returned
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	// The contact id should be returned
	if res.ContactId != "123" {
		t.Error("Unexpected contact id:", res.ContactId)
	}
}
