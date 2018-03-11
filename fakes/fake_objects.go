package fakes

import "github.com/rcw5/vrops-cli/models"

var FakeResources = []models.Resource{
	models.Resource{
		Description: "Description",
		ResourceKey: models.ResourceKey{
			Name:            "my-resource",
			AdapterKindKey:  "my-adapterkind",
			ResourceKindKey: "my-resourcekind",
			ResourceIdentifiers: []models.ResourceIdentifier{
				models.ResourceIdentifier{
					IdentifierType: models.ResourceIdentifierType{
						Name:               "identifier-type",
						DataType:           "string",
						IsPartOfUniqueness: true,
					},
					Value: "value",
				},
			},
		},
		ResourceHealth: "GREEN",
		Identifier:     "an-identifier",
	},
}

var FakeAdapterKinds = []models.AdapterKind{
	models.AdapterKind{
		Key:             "Adapter Key",
		Description:     "Nice long description here",
		AdapterKindType: "Type",
		Name:            "An Adapter",
	},
	models.AdapterKind{
		Key:             "Adapter Key 2",
		Description:     "Nice long description here",
		AdapterKindType: "Type",
		Name:            "An Adapter",
	},
}
