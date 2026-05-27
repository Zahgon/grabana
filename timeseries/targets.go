package timeseries

import (
	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/loki"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
)

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

// WithLokiTarget adds a loki query to the graph.
func WithLokiTarget(query string, options ...loki.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
