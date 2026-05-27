package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/grabana/timeseries/axis"
	"github.com/K-Phoen/grabana/timeseries/fields"
)

var ErrInvalidGradientMode = fmt.Errorf("invalid gradient mode")
var ErrInvalidLineInterpolationMode = fmt.Errorf("invalid line interpolation mode")
var ErrInvalidTooltipMode = fmt.Errorf("invalid tooltip mode")
var ErrInvalidStackMode = fmt.Errorf("invalid stack mode")
var ErrInvalidAxisDisplay = fmt.Errorf("invalid axis display")
var ErrInvalidAxisScale = fmt.Errorf("invalid axis scale")
var ErrInvalidOverrideMatcher = fmt.Errorf("invalid override matcher")

type DashboardTimeSeries struct {
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
	Legend          []string                 `yaml:",omitempty,flow"`
	Alert           *Alert                   `yaml:",omitempty"`
	Visualization   *TimeSeriesVisualization `yaml:",omitempty"`
	Axis            *TimeSeriesAxis          `yaml:",omitempty"`
	Overrides       []TimeSeriesOverride     `yaml:",omitempty"`
}

func (timeseriesPanel DashboardTimeSeries) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}

func (timeseriesPanel DashboardTimeSeries) legend() ([]timeseries.LegendOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (timeseriesPanel DashboardTimeSeries) target(t Target) (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

type TimeSeriesVisualization struct {
	GradientMode      string `yaml:"gradient_mode,omitempty"`
	Tooltip           string `yaml:"tooltip,omitempty"`
	Stack             string `yaml:"stack,omitempty"`
	FillOpacity       *int   `yaml:"fill_opacity,omitempty"`
	PointSize         *int   `yaml:"point_size,omitempty"`
	LineInterpolation string `yaml:"line_interpolation,omitempty"`
	LineWidth         *int   `yaml:"line_width,omitempty"`
	// TODO: draw: {bars: {}, lines: {}}
}

func (timeseriesViz *TimeSeriesVisualization) toOptions() ([]timeseries.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (timeseriesViz *TimeSeriesVisualization) lineInterpolationOption() (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

func (timeseriesViz *TimeSeriesVisualization) gradientModeOption() (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

func (timeseriesViz *TimeSeriesVisualization) tooltipOption() (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

func (timeseriesViz *TimeSeriesVisualization) stackOption() (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

type TimeSeriesAxis struct {
	SoftMin *int     `yaml:"soft_min,omitempty"`
	SoftMax *int     `yaml:"soft_max,omitempty"`
	Min     *float64 `yaml:",omitempty"`
	Max     *float64 `yaml:",omitempty"`

	Decimals *int `yaml:",omitempty"`

	Display string `yaml:",omitempty"`
	Scale   string `yaml:",omitempty"`

	Unit  string `yaml:",omitempty"`
	Label string `yaml:",omitempty"`
}

func (tsAxis *TimeSeriesAxis) toOptions() ([]axis.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tsAxis *TimeSeriesAxis) placementOption() (axis.Option, error) {
	_ = "STUB: not implemented"
	return *new(axis.Option), nil
}

func axisPlacementFromString(input string) (axis.PlacementMode, error) {
	_ = "STUB: not implemented"
	return *new(axis.PlacementMode), nil
}

func (tsAxis *TimeSeriesAxis) scaleOption() (axis.Option, error) {
	_ = "STUB: not implemented"
	return *new(axis.Option), nil
}

type TimeSeriesOverride struct {
	Matcher    TimeSeriesOverrideMatcher `yaml:"match,flow"`
	Properties TimeSeriesOverrideProperties
}

func (override TimeSeriesOverride) toOption() (timeseries.Option, error) {
	_ = "STUB: not implemented"
	return *new(timeseries.Option), nil
}

type TimeSeriesOverrideMatcher struct {
	FieldName *string `yaml:"field_name,omitempty"`
	QueryRef  *string `yaml:"query_ref,omitempty"`
	Regex     *string `yaml:"regex,omitempty"`
	Type      *string `yaml:"field_type,omitempty"`
}

func (matcher TimeSeriesOverrideMatcher) toOption() (fields.Matcher, error) {
	_ = "STUB: not implemented"
	return *new(fields.Matcher), nil
}

type TimeSeriesOverrideProperties struct {
	Unit        *string `yaml:",omitempty"`
	Color       *string `yaml:"color,omitempty"`
	FillOpacity *int    `yaml:"fill_opacity,omitempty"`
	NegativeY   *bool   `yaml:"negative_Y,omitempty"`
	AxisDisplay *string `yaml:"axis_display,omitempty"`
	Stack       *string `yaml:",omitempty"`
}

func (properties TimeSeriesOverrideProperties) toOptions() ([]fields.OverrideOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
