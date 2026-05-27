package singlestat

import (
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a single stat panel.
type Option func(stat *SingleStat) error

// StatType let you set the function that your entire query is reduced into a
// single value with.
type StatType string

const (
	// Min will return the smallest value in the series.
	Min StatType = "min"

	// Max will return the largest value in the series.
	Max StatType = "max"

	// Avg will return the average of all the non-null values in the series.
	Avg StatType = "avg"

	// Current will return the last value in the series. If the series ends on
	// null the previous value will be used.
	Current StatType = "current"

	// Total will return the sum of all the non-null values in the series.
	Total StatType = "total"

	// First will return the first value in the series.
	First StatType = "first"

	// Delta will return the total incremental increase (of a counter) in the
	// series. An attempt is made to account for counter resets, but this will
	// only be accurate for single instance metrics. Used to show total
	// counter increase in time series.
	Delta StatType = "delta"

	// Diff will return difference between ‘current’ (last value) and ‘first’..
	Diff StatType = "diff"

	// Range will return the difference between ‘min’ and ‘max’. Useful to
	// show the range of change for a gauge..
	Range StatType = "range"

	// Name will return the name value in the series.
	Name StatType = "name"
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

// nolint: gochecknoglobals
var valueToTextMapping = 1

// nolint: gochecknoglobals
var rangeToTextMapping = 2

// SingleStat represents a single stat panel.
type SingleStat struct {
	Builder *sdk.Panel
}

// New creates a new single stat panel.
func New(title string, options ...Option) (*SingleStat, error) {
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

// SparkLine displays the spark line summary of the series in addition to the
// single stat.
func SparkLine() Option { _ = "STUB: not implemented"; return *new(Option) }

// FullSparkLine displays a full height spark line summary of the series in
// addition to the single stat.
func FullSparkLine() Option { _ = "STUB: not implemented"; return *new(Option) }

// SparkLineColor sets the line color of the spark line.
func SparkLineColor(color string) Option { _ = "STUB: not implemented"; return *new(Option) }

// SparkLineFillColor sets the color the spark line will be filled with.
func SparkLineFillColor(color string) Option { _ = "STUB: not implemented"; return *new(Option) }

// SparkLineYMin defines the smallest value expected on the Y axis of the spark line.
func SparkLineYMin(value float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// SparkLineYMax defines the largest value expected on the Y axis of the spark line.
func SparkLineYMax(value float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValueType configures how the series will be reduced to a single value.
func ValueType(valueType StatType) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValueFontSize sets the font size used to display the value (eg: "100%").
func ValueFontSize(size string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Prefix sets the text used as prefix of the value.
func Prefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// PrefixFontSize sets the size used for the prefix text (eg: "110%").
func PrefixFontSize(size string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Postfix sets the text used as postfix of the value.
func Postfix(postfix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// PostfixFontSize sets the size used for the postfix text (eg: "110%")
func PostfixFontSize(size string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ColorValue will show the threshold's colors on the value itself.
func ColorValue() Option { _ = "STUB: not implemented"; return *new(Option) }

// ColorBackground will show the threshold's colors in the background.
func ColorBackground() Option { _ = "STUB: not implemented"; return *new(Option) }

// Thresholds change the background and value colors dynamically within the
// panel, depending on the Singlestat value. The threshold is defined by 2
// values which represent 3 ranges that correspond to the three colors directly
// to the right.
func Thresholds(values [2]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Colors define which colors will be applied to the single value based on the
// threshold levels.
func Colors(values [3]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValuesToText allows to translate the value of the summary stat into explicit
// text.
func ValuesToText(mapping []ValueMap) Option { _ = "STUB: not implemented"; return *new(Option) }

// RangesToText allows to translate the value of the summary stat into explicit
// text.
func RangesToText(mapping []RangeMap) Option { _ = "STUB: not implemented"; return *new(Option) }

// Repeat configures repeating a panel for a variable
func Repeat(repeat string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatDirection configures repeating vertical or horizontal
func RepeatDirection(direction sdk.RepeatDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
