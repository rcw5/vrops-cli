package fakes

import "github.com/rcw5/vrops-cli/models"

var FakeResources = models.Resources{
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

var FakeAdapterKinds = models.AdapterKinds{
	models.AdapterKind{
		Key:             "Adapter Key",
		Description:     "Nice long description here",
		AdapterKindType: "Type",
		Name:            "my-adapterkind",
	},
	models.AdapterKind{
		Key:             "Adapter Key 2",
		Description:     "Nice long description here",
		AdapterKindType: "Type",
		Name:            "my-adapterkind-2",
	},
}

var FakeStats = []models.Stat{
	models.Stat{
		StatKey:    "a|stat|key",
		Timestamps: []int64{1620688924000, 1720688924000, 1820688924000},
		Data:       []float64{100, 0, 0},
	},
	models.Stat{
		StatKey:    "another|stat|key",
		Timestamps: []int64{1620688924000, 1720688924000, 1820688924000},
		Data:       []float64{100, 90, 500},
	},
}
