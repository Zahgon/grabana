package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
)

var ErrInvalidColoringTarget = fmt.Errorf("invalid coloring target")
var ErrInvalidSparkLineMode = fmt.Errorf("invalid sparkline mode")
var ErrInvalidSingleStatValueType = fmt.Errorf("invalid single stat value type")

type DashboardSingleStat struct {
	Title           string
	Description     string              `yaml:",omitempty"`
	Span            float32             `yaml:",omitempty"`
	Height          string              `yaml:",omitempty"`
	Transparent     bool                `yaml:",omitempty"`
	Datasource      string              `yaml:",omitempty"`
	Repeat          string              `yaml:",omitempty"`
	RepeatDirection string              `yaml:"repeat_direction,omitempty"`
	Links           DashboardPanelLinks `yaml:",omitempty"`
	Unit            string
	Decimals        *int   `yaml:",omitempty"`
	ValueType       string `yaml:"value_type"`
	ValueFontSize   string `yaml:"value_font_size,omitempty"`
	PrefixFontSize  string `yaml:"prefix_font_size,omitempty"`
	PostfixFontSize string `yaml:"postfix_font_size,omitempty"`
	SparkLine       string `yaml:"sparkline"`
	Targets         []Target
	Thresholds      [2]string
	Colors          [3]string
	Color           []string              `yaml:",omitempty"`
	RangesToText    []singlestat.RangeMap `yaml:"ranges_to_text,omitempty"`
}

func (singleStatPanel DashboardSingleStat) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (singleStatPanel DashboardSingleStat) valueType() (singlestat.Option, error) {
	_ = "STUB: not implemented"
	return *new(singlestat.Option), nil
}

func (singleStatPanel DashboardSingleStat) target(t Target) (singlestat.Option, error) {
	_ = "STUB: not implemented"
	return *new(singlestat.Option), nil
}
