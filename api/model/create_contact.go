package model

import "encoding/json"

type CreateContactData struct {
	FirstName    *string                `json:"firstName"`
	LastName     *string                `json:"lastName"`
	Email        string                 `json:"email"`
	Source       *string                `json:"source"`
	Subscribed   *bool                  `json:"subscribed"`
	UserGroup    *string                `json:"userGroup"`
	UserId       *string                `json:"userId"`
	CustomFields map[string]interface{} `json:"-"`
}

type CreateContactSuccess struct {
	ContactId string
}

func (c CreateContactData) MarshalJSON() ([]byte, error) {
	// step 1: marshall the object
	// this will take care of the fields that are not custom fields
	type dat CreateContactData

	structJson, err := json.Marshal(dat(c)) // use local type to avoid infinite recursion
	if err != nil {
		return nil, err
	}

	// if there's no custom fields, we're done
	if len(c.CustomFields) == 0 {
		return structJson, nil
	}

	// step 2: marshal custom fields
	custom, err := json.Marshal(c.CustomFields)
	if err != nil {
		return nil, err
	}

	// step 3: merge the two json objects
	// first, remove the trailing } from the struct json
	structJson = structJson[:len(structJson)-1]

	// then, remove the leading { from the custom json
	custom = custom[1:]

	// finally, merge the two
	structJson = append(structJson, ',')
	structJson = append(structJson, custom...)

	return structJson, nil
}
