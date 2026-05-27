package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/heatmap"
	"github.com/K-Phoen/grabana/heatmap/axis"
	"github.com/K-Phoen/grabana/row"
)

var ErrInvalidDataFormat = fmt.Errorf("invalid data format")

// DashboardHeatmap represents a heatmap panel.
type DashboardHeatmap struct {
	Title           string
	Description     string              `yaml:",omitempty"`
	Span            float32             `yaml:",omitempty"`
	Height          string              `yaml:",omitempty"`
	Transparent     bool                `yaml:",omitempty"`
	Datasource      string              `yaml:",omitempty"`
	Repeat          string              `yaml:",omitempty"`
	RepeatDirection string              `yaml:"repeat_direction,omitempty"`
	DataFormat      string              `yaml:"data_format,omitempty"`
	HideZeroBuckets bool                `yaml:"hide_zero_buckets"`
	HighlightCards  bool                `yaml:"highlight_cards"`
	Links           DashboardPanelLinks `yaml:",omitempty"`
	Targets         []Target
	ReverseYBuckets bool            `yaml:"reverse_y_buckets,omitempty"`
	Tooltip         *HeatmapTooltip `yaml:",omitempty"`
	YAxis           *HeatmapYAxis   `yaml:",omitempty"`
}

type HeatmapTooltip struct {
	Show          bool
	ShowHistogram bool
	Decimals      *int `yaml:",omitempty"`
}

func (tooltip *HeatmapTooltip) toOptions() []heatmap.Option { _ = "STUB: not implemented"; return nil }

type HeatmapYAxis struct {
	Decimals *int     `yaml:"decimals,omitempty"`
	Unit     string   `yaml:"unit,omitempty"`
	Max      *float64 `yaml:"max,omitempty"`
	Min      *float64 `yaml:"min,omitempty"`
}

func (yaxis *HeatmapYAxis) toOptions() []axis.Option { _ = "STUB: not implemented"; return nil }

func (heatmapPanel DashboardHeatmap) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (heatmapPanel DashboardHeatmap) target(t Target) (heatmap.Option, error) {
	_ = "STUB: not implemented"
	return *new(heatmap.Option), nil
}
