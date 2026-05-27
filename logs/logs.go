package logs

import (
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/grabana/target/loki"
	"github.com/K-Phoen/sdk"
)

// DedupStrategy represents a deduplication strategy.
type DedupStrategy string

const (
	None      DedupStrategy = "none"
	Exact     DedupStrategy = "exact"
	Numbers   DedupStrategy = "numbers"
	Signature DedupStrategy = "signature"
)

// SortOrder represents a sort order.
type SortOrder string

const (
	Asc  SortOrder = "Ascending"
	Desc SortOrder = "Descending"
)

// Option represents an option that can be used to configure a logs panel.
type Option func(logs *Logs) error

// Logs represents a logs panel.
type Logs struct {
	Builder *sdk.Panel
}

// New creates a new logs panel.
func New(title string, options ...Option) (*Logs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// DataSource sets the data source to be used by the panel.
func DataSource(source string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLokiTarget adds a loki query to the graph.
func WithLokiTarget(query string, options ...loki.Option) Option {
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

// Repeat configures repeating a panel for a variable
func Repeat(repeat string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RepeatDirection configures repeating vertical or horizontal
func RepeatDirection(direction sdk.RepeatDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Time displays the "time" column. This is the timestamp associated with the
// log line as reported from the data source.
func Time() Option { _ = "STUB: not implemented"; return *new(Option) }

// UniqueLabels displays the "unique labels" column, which shows only non-common labels.
func UniqueLabels() Option { _ = "STUB: not implemented"; return *new(Option) }

// CommonLabels displays the "common labels".
func CommonLabels() Option { _ = "STUB: not implemented"; return *new(Option) }

// WrapLines enables line wrapping.
func WrapLines() Option { _ = "STUB: not implemented"; return *new(Option) }

// PrettifyJSON pretty prints all JSON logs. This setting does not affect logs
// in any format other than JSON.
func PrettifyJSON() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideLogDetails disables the log details view for each log row.
func HideLogDetails() Option { _ = "STUB: not implemented"; return *new(Option) }

// Order display results in descending or ascending time order.
// The default is Descending, showing the newest logs first.
// Set to Ascending to show the oldest log lines first.
func Order(order SortOrder) Option { _ = "STUB: not implemented"; return *new(Option) }

// Deduplication sets the deduplication strategy.
func Deduplication(dedup DedupStrategy) Option { _ = "STUB: not implemented"; return *new(Option) }
