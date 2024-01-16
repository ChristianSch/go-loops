package main

import (
	"fmt"

	loops "github.com/ChristianSch/go-loops/api"
	"github.com/ChristianSch/go-loops/api/model"
)

func main() {
	client := loops.NewApiClient("https://app.loops.so/api/v1", "your-api-key")
	contact, err := client.ContactAPI().CreateContact(model.CreateContactData{
		Email: "test@example.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("contact id:", contact.ContactId)
}
