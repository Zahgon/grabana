package loki

import (
	"time"

	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a loki query.
type Option func(query *Loki)

// Loki represents a loki query.
type Loki struct {
	Builder sdk.AlertQuery
}

// New creates a new loki query.
func New(ref string, query string, options ...Option) *Loki { _ = "STUB: not implemented"; return nil }

func defaults() []Option { _ = "STUB: not implemented"; return nil }

// TimeRange sets the legend format.
func TimeRange(from time.Duration, to time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Legend sets the legend format.
func Legend(legend string) Option { _ = "STUB: not implemented"; return *new(Option) }
