package gauge

import (
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/scheme"
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a stat panel.
type Option func(gauge *Gauge) error

type ThresholdStep struct {
	Color string
	Value *float64
}

// OrientationMode controls the layout.
type OrientationMode string

const (
	OrientationAuto       OrientationMode = ""
	OrientationHorizontal OrientationMode = "horizontal"
	OrientationVertical   OrientationMode = "vertical"
)

// ReductionType lets you set the function that your entire query is reduced into a
// single value with.
type ReductionType int

const (
	// Min displays the smallest value of the series.
	Min ReductionType = iota
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

// ValueMap allows to map a value into explicit text.
type ValueMap struct {
	Value string
	Text  string
}

// RangeMap allows to map a range of values into explicit text.
type RangeMap struct {
	From string
	To   string
	Text string
}

// Gauge represents a stat panel.
type Gauge struct {
	Builder *sdk.Panel
}

// New creates a new gauge panel.
func New(title string, options ...Option) (*Gauge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// DataSource sets the data source to be used by the panel.
func DataSource(source string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrometheusTarget adds a prometheus query to the graph.
func WithPrometheusTarget(query string, options ...prometheus.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphiteTarget adds a Graphite target to the graph.
func WithGraphiteTarget(query string, options ...graphite.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInfluxDBTarget adds an InfluxDB target to the graph.
func WithInfluxDBTarget(query string, options ...influxdb.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStackdriverTarget adds a stackdriver query to the graph.
func WithStackdriverTarget(target *stackdriver.Stackdriver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Span sets the width of the panel, in grid units. Should be a positive
// number between 1 and 12. Example: 6.
func Span(span float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the panel, in pixels. Example: "400px".
func Height(height string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description annotates the current visualization with a human-readable description.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Transparent makes the background transparent.
func Transparent() Option { _ = "STUB: not implemented"; return *new(Option) }

// Unit sets the unit of the data displayed on this axis.
func Unit(unit string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Decimals sets the number of decimals that should be displayed.
func Decimals(count int) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValueType configures how the series will be reduced to a single value.
func ValueType(valueType ReductionType) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValueFontSize sets the font size used to display the value.
func ValueFontSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// TitleFontSize sets the font size used to display the title.
func TitleFontSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// AbsoluteThresholds changes the background and value colors dynamically within the
// panel, depending on the value. The threshold is defined by a series of steps
// values which, each having a value and an associated color.
func AbsoluteThresholds(steps []ThresholdStep) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// RelativeThresholds changes the background and value colors dynamically within the
// panel, depending on the value. The threshold is defined by a series of steps
// values which, each having a value defined as a percentage and an associated color.
func RelativeThresholds(steps []ThresholdStep) Option {
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

// ColorScheme configures the color scheme.
func ColorScheme(options ...scheme.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// NoValue defines what to show when there is no value.
func NoValue(text string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Orientation changes the orientation of the layout.
func Orientation(mode OrientationMode) Option { _ = "STUB: not implemented"; return *new(Option) }
