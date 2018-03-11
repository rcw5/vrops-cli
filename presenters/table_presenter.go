package presenters

import (
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/rcw5/vrops-cli/models"
)

type TablePresenter struct {
	Buffer io.Writer
}

func (t TablePresenter) PresentAdapterKinds(adapterKinds models.AdapterKinds) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Key", "Name", "Description", "AdapterKindType"})
	for _, adapter := range adapterKinds.AdapterKind {
		table.Append([]string{adapter.Key, adapter.Name, adapter.Description, adapter.AdapterKindType})
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
