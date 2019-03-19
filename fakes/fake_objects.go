package fakes

import "github.com/topflight-technology/vrops-cli/models"

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

var FakeStats = models.Stats{
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

var FakeListStatsResponse = models.ListStatsResponseValuesStatListStats{
	models.ListStatsResponseValuesStatListStat{
		StatKey:      models.ListStatsResponseValuesStatListStatStatKey{Key: "stat|key"},
		Data:         []float64{1, 2, 3, 4, 5},
		Timestamps:   []int64{1, 2, 3, 4, 5},
		IntervalUnit: models.IntervalUnit{Quantifier: 1},
	},
	models.ListStatsResponseValuesStatListStat{
		StatKey:      models.ListStatsResponseValuesStatListStatStatKey{Key: "another-stat|key"},
		Data:         []float64{1, 2, 3, 4, 5},
		Timestamps:   []int64{1, 2, 3, 4, 5},
		IntervalUnit: models.IntervalUnit{Quantifier: 1},
	},
}
