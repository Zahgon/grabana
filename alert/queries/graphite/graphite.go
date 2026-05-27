package graphite

import (
	"time"

	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a graphite query.
type Option func(query *Graphite)

// Graphite represents a graphite query.
type Graphite struct {
	Builder sdk.AlertQuery
}

// New creates a new graphite query.
func New(ref string, query string, options ...Option) *Graphite {
	_ = "STUB: not implemented"
	return nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// TimeRange sets the legend format.
func TimeRange(from time.Duration, to time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Legend sets the legend format.
func Legend(legend string) Option { _ = "STUB: not implemented"; return *new(Option) }
