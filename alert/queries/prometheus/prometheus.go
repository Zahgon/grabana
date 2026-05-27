package prometheus

import (
	"time"

	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a prometheus query.
type Option func(query *Prometheus)

// Prometheus represents a prometheus query.
type Prometheus struct {
	Builder sdk.AlertQuery
}

// New creates a new prometheus query.
func New(ref string, query string, options ...Option) *Prometheus {
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
