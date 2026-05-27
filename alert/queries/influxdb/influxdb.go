package influxdb

import (
	"time"

	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a influxdb query.
type Option func(query *InfluxDB)

// InfluxDB represents a influxdb query.
type InfluxDB struct {
	Builder sdk.AlertQuery
}

// New creates a new influxdb query.
func New(ref string, query string, options ...Option) *InfluxDB {
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
