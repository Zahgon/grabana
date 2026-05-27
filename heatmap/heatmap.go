package heatmap

import (
	"github.com/K-Phoen/grabana/heatmap/axis"
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
	"github.com/K-Phoen/sdk"
)

// DataFormatMode represents the data format modes.
type DataFormatMode string

const (
	// Grafana does the bucketing by going through all time series values
	TimeSeriesBuckets DataFormatMode = "tsbuckets"

	// Each time series already represents a Y-Axis bucket.
	TimeSeries DataFormatMode = "timeseries"
)

// LegendOption allows to configure a legend.
type LegendOption uint16

const (
	// Hide keeps the legend from being displayed.
	Hide LegendOption = iota
)

// Option represents an option that can be used to configure a heatmap panel.
type Option func(stat *Heatmap) error

// Heatmap represents a heatmap panel.
type Heatmap struct {
	Builder *sdk.Panel
}

// New creates a new heatmap panel.
func New(title string, options ...Option) (*Heatmap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

func defaultYAxis() Option { _ = "STUB: not implemented"; return *new(Option) }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// DataSource sets the data source to be used by the panel.
func DataSource(source string) Option { _ = "STUB: not implemented"; return *new(Option) }

// DataFormat sets how the data should be interpreted.
func DataFormat(format DataFormatMode) Option { _ = "STUB: not implemented"; return *new(Option) }

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

// Span sets the width of the panel, in grid units. Should be a positive
// number between 1 and 12. Example: 6.
func Span(span float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the panel, in pixels. Example: "400px".
func Height(height string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description annotates the current visualization with a human-readable description.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Transparent makes the background transparent.
func Transparent() Option { _ = "STUB: not implemented"; return *new(Option) }

// Legend defines what should be shown in the legend.
func Legend(opts ...LegendOption) Option { _ = "STUB: not implemented"; return *new(Option) }

// ShowZeroBuckets forces the display of "zero" buckets.
func ShowZeroBuckets() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideZeroBuckets hides "zero" buckets.
func HideZeroBuckets() Option { _ = "STUB: not implemented"; return *new(Option) }

// HighlightCards highlights bucket cards.
func HighlightCards() Option { _ = "STUB: not implemented"; return *new(Option) }

// NoHighlightCards disables the highlighting of bucket cards.
func NoHighlightCards() Option { _ = "STUB: not implemented"; return *new(Option) }

// ReverseYBuckets reverses the order of bucket on the Y-axis.
func ReverseYBuckets() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideTooltip prevents the tooltip from being displayed.
func HideTooltip() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideTooltipHistogram prevents the histograms from being displayed in tooltips.
// Histogram represents the distribution of the bucket values for the specific timestamp.
func HideTooltipHistogram() Option { _ = "STUB: not implemented"; return *new(Option) }

// TooltipDecimals sets the number of decimals to be displayed in tooltips.
func TooltipDecimals(decimals int) Option { _ = "STUB: not implemented"; return *new(Option) }

// HideXAxis prevents the X-axis from being displayed.
func HideXAxis() Option { _ = "STUB: not implemented"; return *new(Option) }

// YAxis configures the Y axis.
func YAxis(opts ...axis.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// Repeat configures repeating a panel for a variable
func Repeat(repeat string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatDirection configures repeating vertical or horizontal
func RepeatDirection(direction sdk.RepeatDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
