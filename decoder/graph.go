package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/axis"
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/row"
)

var ErrInvalidLegendAttribute = fmt.Errorf("invalid legend attribute")

type DashboardGraph struct {
	Title           string
	Description     string  `yaml:",omitempty"`
	Span            float32 `yaml:",omitempty"`
	Height          string  `yaml:",omitempty"`
	Transparent     bool    `yaml:",omitempty"`
	Datasource      string  `yaml:",omitempty"`
	Repeat          string  `yaml:",omitempty"`
	RepeatDirection string  `yaml:"repeat_direction,omitempty"`
	Targets         []Target
	Links           DashboardPanelLinks `yaml:",omitempty"`
	Axes            *GraphAxes          `yaml:",omitempty"`
	Legend          []string            `yaml:",omitempty,flow"`
	Alert           *Alert              `yaml:",omitempty"`
	Visualization   *GraphVisualization `yaml:",omitempty"`
}

func (graphPanel DashboardGraph) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

type GraphSeriesOverride struct {
	Alias     string
	Color     string `yaml:",omitempty"`
	Dashes    *bool  `yaml:",omitempty"`
	Lines     *bool  `yaml:",omitempty"`
	Fill      *int   `yaml:",omitempty"`
	LineWidth *int   `yaml:"line_width,omitempty"`
}

func (override *GraphSeriesOverride) toOption() graph.Option {
	_ = "STUB: not implemented"
	return *new(graph.Option)
}

type GraphVisualization struct {
	NullValue string                `yaml:",omitempty"`
	Staircase bool                  `yaml:",omitempty"`
	Overrides []GraphSeriesOverride `yaml:"overrides,omitempty"`
}

func (graphViz *GraphVisualization) toOptions() []graph.Option {
	_ = "STUB: not implemented"
	return nil
}

func (graphPanel *DashboardGraph) legend() ([]graph.LegendOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (graphPanel *DashboardGraph) target(t Target) (graph.Option, error) {
	_ = "STUB: not implemented"
	return *new(graph.Option), nil
}

type GraphAxis struct {
	Hidden  *bool    `yaml:",omitempty"`
	Label   string   `yaml:",omitempty"`
	Unit    *string  `yaml:",omitempty"`
	Min     *float64 `yaml:",omitempty"`
	Max     *float64 `yaml:",omitempty"`
	LogBase int      `yaml:"log_base"`
}

func (a GraphAxis) toOptions() []axis.Option { _ = "STUB: not implemented"; return nil }

type GraphAxes struct {
	Left   *GraphAxis `yaml:",omitempty"`
	Right  *GraphAxis `yaml:",omitempty"`
	Bottom *GraphAxis `yaml:",omitempty"`
}
