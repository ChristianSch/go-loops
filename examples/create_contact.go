package main

import (
	"fmt"

	loops "github.com/graileanu/go-loops/api"
	"github.com/graileanu/go-loops/api/model"
)

func main() {
	// create a new client
	client := loops.NewApiClient("https://app.loops.so/api/v1", "your-api-key")

	// create a new contact with only the mandatory email field
	// check model.CreateContactData for all available fields
	contact, err := client.ContactAPI().CreateContact(model.CreateContactData{
		Email: "test@example.com",
		// you can pass arbitrary custom fields
		CustomFields: map[string]interface{}{
			"test":     "foobar",
			"customer": true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("contact id:", contact.ContactId)
}
