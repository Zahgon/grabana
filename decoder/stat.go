package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/stat"
)

var ErrInvalidStatOrientation = fmt.Errorf("invalid orientation")
var ErrInvalidStatColorMode = fmt.Errorf("invalid stat color mode")
var ErrInvalidStatThresholdMode = fmt.Errorf("invalid stat threshold mode")
var ErrInvalidStatTextMode = fmt.Errorf("invalid text mode")
var ErrInvalidStatValueType = fmt.Errorf("invalid stat value type")

type StatThresholdStep struct {
	Color string
	Value *float64 `yaml:",omitempty"`
}

type DashboardStat struct {
	Title           string
	Description     string              `yaml:",omitempty"`
	Span            float32             `yaml:",omitempty"`
	Height          string              `yaml:",omitempty"`
	Transparent     bool                `yaml:",omitempty"`
	Datasource      string              `yaml:",omitempty"`
	Repeat          string              `yaml:",omitempty"`
	RepeatDirection string              `yaml:"repeat_direction,omitempty"`
	Links           DashboardPanelLinks `yaml:",omitempty"`
	Targets         []Target

	Unit     string `yaml:",omitempty"`
	Decimals *int   `yaml:",omitempty"`

	SparkLine     bool   `yaml:"sparkline,omitempty"`
	Orientation   string `yaml:",omitempty"`
	Text          string `yaml:",omitempty"`
	ValueType     string `yaml:"value_type,omitempty"`
	ColorMode     string `yaml:"color_mode,omitempty"`
	TitleFontSize int    `yaml:"title_font_size,omitempty"`
	ValueFontSize int    `yaml:"value_font_size,omitempty"`

	ThresholdMode string              `yaml:"threshold_mode,omitempty"`
	Thresholds    []StatThresholdStep `yaml:",omitempty"`
}

func (statPanel DashboardStat) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (statPanel DashboardStat) thresholds() (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}

func (statPanel DashboardStat) colorMode() (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}

func (statPanel DashboardStat) valueType() (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}

func (statPanel DashboardStat) orientationOpt() (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}

func (statPanel DashboardStat) textOpt() (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}

func (statPanel DashboardStat) target(t Target) (stat.Option, error) {
	_ = "STUB: not implemented"
	return *new(stat.Option), nil
}
