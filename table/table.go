package table

import (
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a table panel.
type Option func(table *Table) error

// AggregationType represents an aggregation function used on values returned
// by the query.
type AggregationType string

const (
	// AVG aggregates results by computing the average.
	AVG AggregationType = "avg"

	// Count aggregates results by counting them.
	Count AggregationType = "count"

	// Current aggregates results by keeping only the current value.
	Current AggregationType = "current"

	// Min aggregates results by keeping only the smallest value.
	Min AggregationType = "min"

	// Max aggregates results by keeping only the largest value.
	Max AggregationType = "max"
)

// Aggregation configures how to display an aggregate in the table.
type Aggregation struct {
	Label string
	Type  AggregationType
}

// Table represents a table panel.
type Table struct {
	Builder *sdk.Panel
}

// New creates a new table panel.
func New(title string, options ...Option) (*Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrometheusTarget adds a prometheus target to the table.
func WithPrometheusTarget(query string, options ...prometheus.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphiteTarget adds a Graphite target to the table.
func WithGraphiteTarget(query string, options ...graphite.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInfluxDBTarget adds an InfluxDB target to the table.
func WithInfluxDBTarget(query string, options ...influxdb.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// HideColumn hides the column having a label matching the given pattern.
func HideColumn(columnLabelPattern string) Option { _ = "STUB: not implemented"; return *new(Option) }

// TimeSeriesToRows displays the data in rows.
func TimeSeriesToRows() Option { _ = "STUB: not implemented"; return *new(Option) }

// TimeSeriesToColumns displays the data in columns.
func TimeSeriesToColumns() Option { _ = "STUB: not implemented"; return *new(Option) }

// AsJSON displays the data as JSON.
func AsJSON() Option { _ = "STUB: not implemented"; return *new(Option) }

// AsTable displays the data as a table.
func AsTable() Option { _ = "STUB: not implemented"; return *new(Option) }

// AsAnnotations displays the data as annotations.
func AsAnnotations() Option { _ = "STUB: not implemented"; return *new(Option) }

// AsTimeSeriesAggregations displays the data according to the given aggregation methods.
func AsTimeSeriesAggregations(aggregations []Aggregation) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// DataSource sets the data source to be used by the table.
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
