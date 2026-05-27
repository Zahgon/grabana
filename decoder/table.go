package decoder

import (
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/table"
)

// DashboardTable represents a table panel.
type DashboardTable struct {
	Title                  string
	Description            string              `yaml:",omitempty"`
	Span                   float32             `yaml:",omitempty"`
	Height                 string              `yaml:",omitempty"`
	Transparent            bool                `yaml:",omitempty"`
	Datasource             string              `yaml:",omitempty"`
	Links                  DashboardPanelLinks `yaml:",omitempty"`
	Targets                []Target
	HiddenColumns          []string            `yaml:"hidden_columns,flow"`
	TimeSeriesAggregations []table.Aggregation `yaml:"time_series_aggregations"`
}

func (tablePanel DashboardTable) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (tablePanel *DashboardTable) target(t Target) (table.Option, error) {
	_ = "STUB: not implemented"
	return *new(table.Option), nil
}
