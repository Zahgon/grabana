package timeseries

import (
	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/scheme"
	"github.com/K-Phoen/grabana/timeseries/axis"
	"github.com/K-Phoen/grabana/timeseries/fields"
	"github.com/K-Phoen/grabana/timeseries/threshold"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a graph panel.
type Option func(timeseries *TimeSeries) error

// TooltipMode configures which series will be displayed in the tooltip.
type TooltipMode string

const (
	// SingleSeries will only display the hovered series.
	SingleSeries TooltipMode = "single"
	// AllSeries will display all series.
	AllSeries TooltipMode = "multi"
	// NoSeries will hide the tooltip completely.
	NoSeries TooltipMode = "none"
)

// StackMode configures mode of series stacking.
type StackMode string

const (
	// Unstacked will not stack series
	Unstacked StackMode = "none"
	// NormalStack will stack series as absolute numbers
	NormalStack StackMode = "normal"
	// PercentStack will stack series as percents
	PercentStack StackMode = "percent"
)

// LineInterpolationMode defines how Grafana interpolates series lines when drawn as lines.
type LineInterpolationMode string

const (
	// Points are joined by straight lines.
	Linear LineInterpolationMode = "linear"
	// Points are joined by curved lines resulting in smooth transitions between points.
	Smooth LineInterpolationMode = "smooth"
	// The line is displayed as steps between points. Points are rendered at the end of the step.
	StepBefore LineInterpolationMode = "stepBefore"
	// Line is displayed as steps between points. Points are rendered at the beginning of the step.
	StepAfter LineInterpolationMode = "stepAfter"
)

// BarAlignment defines how Grafana aligns bars.
type BarAlignment int

const (
	// The bar is drawn around the point. The point is placed in the center of the bar.
	AlignCenter BarAlignment = 0
	// The bar is drawn before the point. The point is placed on the trailing corner of the bar.
	AlignBefore BarAlignment = -1
	// The bar is drawn after the point. The point is placed on the leading corner of the bar.
	AlignAfter BarAlignment = 1
)

// GradientType defines the mode of the gradient fill.
type GradientType string

const (
	// No gradient fill.
	NoGradient GradientType = "none"
	// Transparency of the gradient is calculated based on the values on the y-axis.
	// Opacity of the fill is increasing with the values on the Y-axis.
	Opacity GradientType = "opacity"
	// Gradient color is generated based on the hue of the line color.
	Hue GradientType = "hue"
	// In this mode the whole bar will use a color gradient defined by the color scheme.
	Scheme GradientType = "scheme"
)

// LegendOption allows to configure a legend.
type LegendOption uint16

const (
	// Hide keeps the legend from being displayed.
	Hide LegendOption = iota
	// AsTable displays the legend as a table.
	AsTable
	// AsList displays the legend as a list.
	AsList
	// Bottom displays the legend below the graph.
	Bottom
	// ToTheRight displays the legend on the right side of the graph.
	ToTheRight

	// Min displays the smallest value of the series.
	Min
	// Max displays the largest value of the series.
	Max
	// Avg displays the average of the series.
	Avg

	// First displays the first value of the series.
	First
	// FirstNonNull displays the first non-null value of the series.
	FirstNonNull
	// Last displays the last value of the series.
	Last
	// LastNonNull displays the last non-null value of the series.
	LastNonNull

	// Total displays the sum of values in the series.
	Total
	// Count displays the number of value in the series.
	Count
	// Range displays the difference between the minimum and maximum values.
	Range
)

// TimeSeries represents a time series panel.
type TimeSeries struct {
	Builder *sdk.Panel
	Alert   *alert.Alert
}

// New creates a new time series panel.
func New(title string, options ...Option) (*TimeSeries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// DataSource sets the data source to be used by the graph.
func DataSource(source string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Tooltip configures the tooltip content.
func Tooltip(mode TooltipMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// LineWidth defines the width of the line for a series (default 1, max 10, 0 is none).
func LineWidth(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Stack defines if the series should be stacked and using which mode (default not stacked).
func Stack(value StackMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// FillOpacity defines the opacity level of the series. The lower the value, the more transparent.
func FillOpacity(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// PointSize adjusts the size of points.
func PointSize(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Lines displays the series as lines, with a given interpolation strategy.
func Lines(mode LineInterpolationMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Bars displays the series as bars, with a given alignment strategy.
func Bars(alignment BarAlignment) Option { _ = "STUB: not implemented"; return *new(Option) }

// Points displays the series as points.
func Points() Option { _ = "STUB: not implemented"; return *new(Option) }

// GradientMode sets the mode of the gradient fill.
func GradientMode(mode GradientType) Option { _ = "STUB: not implemented"; return *new(Option) }

// Axis configures the axis for this time series.
func Axis(options ...axis.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Thresholds configures the thresholds for this time series.
func Thresholds(options ...threshold.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// ColorScheme configures the color scheme.
func ColorScheme(options ...scheme.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Legend defines what should be shown in the legend.
func Legend(opts ...LegendOption) Option { _ = "STUB: not implemented"; return *new(Option) }

// Span sets the width of the panel, in grid units. Should be a positive
// number between 1 and 12. Example: 6.
func Span(span float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the panel, in pixels. Example: "400px".
func Height(height string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description annotates the current visualization with a human-readable description.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Transparent makes the background transparent.
func Transparent() Option { _ = "STUB: not implemented"; return *new(Option) }

// Alert creates an alert for this graph.
func Alert(name string, opts ...alert.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Repeat configures repeating a panel for a variable
func Repeat(repeat string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatDirection configures repeating vertical or horizontal
func RepeatDirection(direction sdk.RepeatDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// FieldOverride allows overriding visualization options.
func FieldOverride(m fields.Matcher, opts ...fields.OverrideOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
