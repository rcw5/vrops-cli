package presenters

import (
	"io"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rcw5/vrops-cli/models"
)

type TablePresenter struct {
	Buffer io.Writer
}

func (t TablePresenter) PresentAdapterKinds(adapterKinds models.AdapterKinds) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Key", "Name", "Description", "AdapterKindType"})
	for _, adapter := range adapterKinds {
		table.Append([]string{adapter.Key, adapter.Name, adapter.Description, adapter.AdapterKindType})
	}
	table.Render()
}

func (t TablePresenter) PresentResources(resources models.Resources) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Name", "Identifier", "AdapterKind", "ResourceKind", "Description", "Health"})
	for _, resource := range resources {
		table.Append([]string{resource.ResourceKey.Name, resource.Identifier, resource.ResourceKey.AdapterKindKey,
			resource.ResourceKey.ResourceKindKey, resource.Description, resource.ResourceHealth})
	}
	table.Render()
}

func (t TablePresenter) PresentStats(stats models.ListStatsResponseValuesStatListStats) {
	table := tablewriter.NewWriter(t.Buffer)
	table.SetHeader([]string{"Name", "Time", "Value"})
	table.SetAutoWrapText(false)
	table.SetAutoMergeCells(false)
	for _, stat := range stats {
		for i := range stat.Data {
			table.Append([]string{stat.StatKey.Key, t.timestampToTime(stat.Timestamps[i]).String(), strconv.FormatFloat(stat.Data[i], 'f', -1, 64)})
		}
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

func (t TablePresenter) timestampToTime(m int64) time.Time {
	return time.Unix(0, m*int64(time.Millisecond))
}
