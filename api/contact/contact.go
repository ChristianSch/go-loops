package contact

import (
	"github.com/ChristianSch/go-loops/api/model"
)

type ContactAPI struct {
	client model.ApiClient
}

func NewContactAPI(client model.ApiClient) model.ContactAPI {
	return ContactAPI{client: client}
}

// enforce interface implementation
var _ model.ContactAPI = ContactAPI{}
