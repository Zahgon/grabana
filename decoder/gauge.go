package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/gauge"
	"github.com/K-Phoen/grabana/row"
)

var ErrInvalidGaugeThresholdMode = fmt.Errorf("invalid gauge threshold mode")
var ErrInvalidGaugeValueType = fmt.Errorf("invalid gauge value type")
var ErrInvalidGaugeOrientation = fmt.Errorf("invalid gauge orientation")

type GaugeThresholdStep struct {
	Color string
	Value *float64 `yaml:",omitempty"`
}

type DashboardGauge struct {
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

	Orientation   string `yaml:",omitempty"`
	ValueType     string `yaml:"value_type,omitempty"`
	TitleFontSize int    `yaml:"title_font_size,omitempty"`
	ValueFontSize int    `yaml:"value_font_size,omitempty"`

	ThresholdMode string               `yaml:"threshold_mode,omitempty"`
	Thresholds    []GaugeThresholdStep `yaml:",omitempty"`
}

func (gaugePanel DashboardGauge) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (gaugePanel DashboardGauge) thresholds() (gauge.Option, error) {
	_ = "STUB: not implemented"
	return *new(gauge.Option), nil
}

func (gaugePanel DashboardGauge) valueType() (gauge.Option, error) {
	_ = "STUB: not implemented"
	return *new(gauge.Option), nil
}

func (gaugePanel DashboardGauge) orientationOpt() (gauge.Option, error) {
	_ = "STUB: not implemented"
	return *new(gauge.Option), nil
}

func (gaugePanel DashboardGauge) target(t Target) (gauge.Option, error) {
	_ = "STUB: not implemented"
	return *new(gauge.Option), nil
}
