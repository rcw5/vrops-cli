package presenters

import (
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/rcw5/vrops-cli/models"
)

type TablePresenter struct {
	Buffer io.Writer
}

func (t TablePresenter) PresentAdapterKinds(adapterKinds []models.AdapterKind) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Key", "Name", "Description", "AdapterKindType"})
	for _, adapter := range adapterKinds {
		table.Append([]string{adapter.Key, adapter.Name, adapter.Description, adapter.AdapterKindType})
	}
	table.Render()
}

func (t TablePresenter) PresentResources(resources []models.Resource) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Name", "Identifier", "AdapterKind", "ResourceKind", "Description", "Health"})
	for _, resource := range resources {
		table.Append([]string{resource.ResourceKey.Name, resource.Identifier, resource.ResourceKey.AdapterKindKey,
			resource.ResourceKey.ResourceKindKey, resource.Description, resource.ResourceHealth})
	}
	table.Render()
}

func (t TablePresenter) PresentResourceKinds(resourceKinds []string) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Name"})
	for _, resource := range resourceKinds {
		table.Append([]string{resource})
	}
	table.Render()
}
