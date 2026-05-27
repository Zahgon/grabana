package row

import (
	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/gauge"
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/heatmap"
	"github.com/K-Phoen/grabana/logs"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/stat"
	"github.com/K-Phoen/grabana/table"
	"github.com/K-Phoen/grabana/text"
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a row.
type Option func(row *Row) error

// Row represents a dashboard row.
type Row struct {
	builder *sdk.Row
	alerts  []*alert.Alert
}

// New creates a new row.
func New(board *sdk.Board, title string, options ...Option) (*Row, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// Alerts returns a list of alerts defined within this row.
func (row *Row) Alerts() []*alert.Alert {
	_ = "STUB: not implemented"

	// WithGraph adds a "graph" panel in the row.
	// Deprecated: use WithTimeSeries() instead.
	return nil
}

func WithGraph(title string, options ...graph.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTimeSeries adds a "timeseries" panel in the row.
func WithTimeSeries(title string, options ...timeseries.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGauge adds a "gauge" panel in the row.
func WithGauge(title string, options ...gauge.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLogs adds a "logs" panel in the row.
func WithLogs(title string, options ...logs.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSingleStat adds a "single stat" panel in the row.
// Deprecated: use WithStat() instead
func WithSingleStat(title string, options ...singlestat.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStat adds a "stat" panel in the row.
func WithStat(title string, options ...stat.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTable adds a "table" panel in the row.
func WithTable(title string, options ...table.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithText adds a "text" panel in the row.
func WithText(title string, options ...text.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHeatmap adds a "heatmap" panel in the row.
func WithHeatmap(title string, options ...heatmap.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ShowTitle ensures that the title of the row will be displayed.
func ShowTitle() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideTitle ensures that the title of the row will NOT be displayed.
func HideTitle() Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatFor will repeat the row for all values of the given variable.
func RepeatFor(variable string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Collapse makes the row collapsed by default.
func Collapse() Option { _ = "STUB: not implemented"; return *new(Option) }
