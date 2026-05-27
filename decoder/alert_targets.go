package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/alert/queries/stackdriver"
)

var ErrMissingRef = fmt.Errorf("target ref missing")
var ErrInvalidLookback = fmt.Errorf("invalid lookback")

type AlertTarget struct {
	Prometheus  *AlertPrometheus  `yaml:",omitempty"`
	Loki        *AlertLoki        `yaml:",omitempty"`
	Graphite    *AlertGraphite    `yaml:",omitempty"`
	Stackdriver *AlertStackdriver `yaml:",omitempty"`
}

func (t AlertTarget) toOption() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

type AlertPrometheus struct {
	Ref      string `yaml:",omitempty"`
	Query    string
	Legend   string `yaml:",omitempty"`
	Lookback string `yaml:",omitempty"`
}

func (t AlertPrometheus) toOptions() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

type AlertLoki struct {
	Ref      string `yaml:",omitempty"`
	Query    string
	Legend   string `yaml:",omitempty"`
	Lookback string `yaml:",omitempty"`
}

func (t AlertLoki) toOptions() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

type AlertGraphite struct {
	Ref      string `yaml:",omitempty"`
	Query    string
	Lookback string `yaml:",omitempty"`
}

func (t AlertGraphite) toOptions() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

type AlertStackdriver struct {
	Ref      string `yaml:",omitempty"`
	Lookback string `yaml:",omitempty"`

	Project      string `yaml:",omitempty"`
	Type         string
	Metric       string
	Filters      StackdriverAlertFilters    `yaml:",omitempty"`
	Aggregation  string                     `yaml:",omitempty"`
	Alignment    *StackdriverAlertAlignment `yaml:",omitempty"`
	Legend       string                     `yaml:",omitempty"`
	Preprocessor string                     `yaml:",omitempty"`
	Hidden       bool                       `yaml:",omitempty"`
	GroupBy      []string                   `yaml:"group_by,omitempty"`
}

type StackdriverAlertFilters struct {
	Eq         map[string]string `yaml:",omitempty"`
	Neq        map[string]string `yaml:",omitempty"`
	Matches    map[string]string `yaml:",omitempty"`
	NotMatches map[string]string `yaml:"not_matches,omitempty"`
}

type StackdriverAlertAlignment struct {
	Method string
	Period string
}

func (t AlertStackdriver) toOptions() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

func (t AlertStackdriver) targetOptions() ([]stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t AlertStackdriver) aggregation() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}

func (t AlertStackdriver) preprocessor() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}

func (filters StackdriverAlertFilters) toOptions() []stackdriver.FilterOption {
	_ = "STUB: not implemented"
	return nil
}

func (t StackdriverAlertAlignment) toOption() (stackdriver.Option, error) {
	_ = "STUB: not implemented"
	return *new(stackdriver.Option), nil
}
