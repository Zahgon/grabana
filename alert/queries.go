package alert

import (
	"github.com/K-Phoen/grabana/alert/queries/graphite"
	"github.com/K-Phoen/grabana/alert/queries/influxdb"
	"github.com/K-Phoen/grabana/alert/queries/loki"
	"github.com/K-Phoen/grabana/alert/queries/prometheus"
	"github.com/K-Phoen/grabana/alert/queries/stackdriver"
)

// WithPrometheusQuery adds a prometheus query to the alert.
func WithPrometheusQuery(ref string, query string, options ...prometheus.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphiteQuery adds a graphite query to the alert.
func WithGraphiteQuery(ref string, query string, options ...graphite.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLokiQuery adds a loki query to the alert.
func WithLokiQuery(ref string, query string, options ...loki.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStackdriverQuery adds a Stackdriver query to the alert.
func WithStackdriverQuery(query *stackdriver.Stackdriver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInfluxDBQuery adds an InfluxDB query to the alert.
func WithInfluxDBQuery(ref string, query string, options ...influxdb.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
