package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/target/graphite"
	"github.com/K-Phoen/grabana/target/influxdb"
	"github.com/K-Phoen/grabana/target/loki"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/target/stackdriver"
)

var ErrTargetNotConfigured = fmt.Errorf("target not configured")
var ErrInvalidStackdriverType = fmt.Errorf("invalid stackdriver target type")
var ErrInvalidStackdriverAggregation = fmt.Errorf("invalid stackdriver aggregation type")
var ErrInvalidStackdriverPreprocessor = fmt.Errorf("invalid stackdriver preprocessor")
var ErrInvalidStackdriverAlignment = fmt.Errorf("invalid stackdriver alignment method")

type Target struct {
	Prometheus  *PrometheusTarget  `yaml:",omitempty"`
	Graphite    *GraphiteTarget    `yaml:",omitempty"`
	InfluxDB    *InfluxDBTarget    `yaml:"influxdb,omitempty"`
	Stackdriver *StackdriverTarget `yaml:",omitempty"`
	Loki        *LokiTarget        `yaml:",omitempty"`
}

type PrometheusTarget struct {
	Query          string
	Legend         string `yaml:",omitempty"`
	Ref            string `yaml:",omitempty"`
	Hidden         bool   `yaml:",omitempty"`
	Format         string `yaml:",omitempty"`
	Instant        bool   `yaml:",omitempty"`
	IntervalFactor *int   `yaml:"interval_factor,omitempty"`
}

func (t PrometheusTarget) toOptions() []prometheus.Option { _ = "STUB: not implemented"; return nil }

type LokiTarget struct {
	Query  string
	Legend string `yaml:",omitempty"`
	Ref    string `yaml:",omitempty"`
	Hidden bool   `yaml:",omitempty"`
}

func (t LokiTarget) toOptions() []loki.Option { _ = "STUB: not implemented"; return nil }

type GraphiteTarget struct {
	Query  string
	Ref    string `yaml:",omitempty"`
	Hidden bool   `yaml:",omitempty"`
}

func (t GraphiteTarget) toOptions() []graphite.Option { _ = "STUB: not implemented"; return nil }

type InfluxDBTarget struct {
	Query  string
	Ref    string `yaml:",omitempty"`
	Hidden bool   `yaml:",omitempty"`
}

func (t InfluxDBTarget) toOptions() []influxdb.Option { _ = "STUB: not implemented"; return nil }

type StackdriverTarget struct {
	Project      string
	Type         string
	Metric       string
	Filters      StackdriverFilters    `yaml:",omitempty"`
	Aggregation  string                `yaml:",omitempty"`
	Alignment    *StackdriverAlignment `yaml:",omitempty"`
	Legend       string                `yaml:",omitempty"`
	Preprocessor string                `yaml:",omitempty"`
	Ref          string                `yaml:",omitempty"`
	Hidden       bool                  `yaml:",omitempty"`
	GroupBy      []string              `yaml:"group_by,omitempty"`
}

type StackdriverFilters struct {
	Eq         map[string]string `yaml:",omitempty"`
	Neq        map[string]string `yaml:",omitempty"`
	Matches    map[string]string `yaml:",omitempty"`
	NotMatches map[string]string `yaml:"not_matches,omitempty"`
}

type StackdriverAlignment struct {
	Method string
	Period string
}

func (t StackdriverTarget) toTarget() (*stackdriver.Stackdriver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t StackdriverTarget) toOptions() ([]stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t StackdriverTarget) aggregation() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}

func (t StackdriverTarget) preprocessor() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}

func (filters StackdriverFilters) toOptions() []stackdriver.FilterOption {
	_ = "STUB: not implemented"
	return nil
}

func (t StackdriverAlignment) toOption() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}
