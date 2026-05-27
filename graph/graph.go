package graph

import (
	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/axis"
	"github.com/K-Phoen/grabana/graph/series"
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a graph panel.
type Option func(graph *Graph) error

// DrawMode represents a type of visualization that will be drawn in the graph
// (lines, bars, points)
type DrawMode uint8

const (
	// Bars will display bars.
	Bars DrawMode = iota
	// Lines will display lines.
	Lines
	// Points will display points.
	Points
)

// NullValue describes how null values are displayed.
type NullValue string

const (
	// AsZero treats null values as zero values.
	AsZero NullValue = "null as zero"

	// AsNull treats null values as null.
	AsNull NullValue = "null"

	// Connected connects null values.
	Connected NullValue = "connected"
)

// LegendOption allows to configure a legend.
type LegendOption uint16

const (
	// Hide keeps the legend from being displayed.
	Hide LegendOption = iota
	// AsTable displays the legend as a table.
	AsTable
	// ToTheRight displays the legend on the right side of the graph.
	ToTheRight
	// Min displays the smallest value of the series.
	Min
	// Max displays the largest value of the series.
	Max
	// Avg displays the average of the series.
	Avg
	// Current displays the current value of the series.
	Current
	// Total displays the total value of the series.
	Total
	// NoNullSeries hides series with only null values from the legend.
	NoNullSeries
	// NoZeroSeries hides series with only 0 values from the legend.
	NoZeroSeries
)

// Graph represents a graph panel.
type Graph struct {
	Builder *sdk.Panel
	Alert   *alert.Alert
}

// New creates a new graph panel.
func New(title string, options ...Option) (*Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

func defaultAxes() Option { _ = "STUB: not implemented"; return *new(Option) }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrometheusTarget adds a prometheus query to the graph.
func WithPrometheusTarget(query string, options ...prometheus.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphiteTarget adds a Graphite target to the table.
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

// DataSource sets the data source to be used by the graph.
func DataSource(source string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Span sets the width of the panel, in grid units. Should be a positive
// number between 1 and 12. Example: 6.
func Span(span float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the panel, in pixels. Example: "400px".
func Height(height string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description annotates the current visualization with a human-readable description.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Transparent makes the background transparent.
func Transparent() Option { _ = "STUB: not implemented"; return *new(Option) }

// LeftYAxis configures the left Y axis.
func LeftYAxis(opts ...axis.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// RightYAxis configures the right Y axis.
func RightYAxis(opts ...axis.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// XAxis configures the X axis.
func XAxis(opts ...axis.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Alert creates an alert for this graph.
func Alert(name string, opts ...alert.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Draw specifies how the graph will be drawn.
func Draw(modes ...DrawMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Fill defines the amount of color fill for a series (default 1, max 10, 0 is none).
func Fill(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// LineWidth defines the width of the line for a series (default 1, max 10, 0 is none).
func LineWidth(value uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// Staircase draws adjacent points as staircase.
func Staircase() Option { _ = "STUB: not implemented"; return *new(Option) }

// PointRadius adjusts the size of points when Points are selected as Draw Mode.
func PointRadius(value float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Null configures how null values are displayed.
func Null(mode NullValue) Option { _ = "STUB: not implemented"; return *new(Option) }

// Repeat configures repeating a panel for a variable
func Repeat(repeat string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatDirection configures repeating vertical or horizontal
func RepeatDirection(direction sdk.RepeatDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SeriesOverride configures how null values are displayed.
// See https://grafana.com/docs/grafana/latest/panels/field-options/
func SeriesOverride(opts ...series.OverrideOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Legend defines what should be shown in the legend.
func Legend(opts ...LegendOption) Option { _ = "STUB: not implemented"; return *new(Option) }
