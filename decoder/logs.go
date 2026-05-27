package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/logs"
	"github.com/K-Phoen/grabana/row"
)

var ErrInvalidSortOrder = fmt.Errorf("invalid sort order")
var ErrInvalidDeduplicationStrategy = fmt.Errorf("invalid deduplication strategy")

type DashboardLogs struct {
	Title           string
	Description     string              `yaml:",omitempty"`
	Span            float32             `yaml:",omitempty"`
	Height          string              `yaml:",omitempty"`
	Transparent     bool                `yaml:",omitempty"`
	Datasource      string              `yaml:",omitempty"`
	Repeat          string              `yaml:",omitempty"`
	RepeatDirection string              `yaml:"repeat_direction,omitempty"`
	Links           DashboardPanelLinks `yaml:",omitempty"`
	Targets         []LogsTarget        `yaml:",omitempty"`
	Visualization   *LogsVisualization  `yaml:",omitempty"`
}

type LogsTarget struct {
	Loki *LokiTarget `yaml:",omitempty"`
}

func (panel DashboardLogs) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (panel DashboardLogs) target(t LogsTarget) (logs.Option, error) {
	_ = "STUB: not implemented"
	return *new(logs.Option), nil
}

type LogsVisualization struct {
	Time           bool   `yaml:",omitempty"`
	UniqueLabels   bool   `yaml:"unique_labels,omitempty"`
	CommonLabels   bool   `yaml:"common_labels,omitempty"`
	WrapLines      bool   `yaml:"wrap_lines,omitempty"`
	PrettifyJSON   bool   `yaml:"prettify_json,omitempty"`
	HideLogDetails bool   `yaml:"hide_log_details,omitempty"`
	Order          string `yaml:",omitempty"`
	Deduplication  string `yaml:",omitempty"`
}

func (viz *LogsVisualization) toOptions() ([]logs.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
