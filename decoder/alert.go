package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/alert"
)

var ErrNoAlertThresholdDefined = fmt.Errorf("no threshold defined")
var ErrInvalidAlertValueFunc = fmt.Errorf("invalid alert value function")
var ErrInvalidAlertOperand = fmt.Errorf("invalid alert operand")
var ErrNoConditionOnAlert = fmt.Errorf("no condition defined on alert")
var ErrNoTargetOnAlert = fmt.Errorf("no target defined on alert")

type Alert struct {
	Summary     string
	Description string            `yaml:",omitempty"`
	Runbook     string            `yaml:",omitempty"`
	Tags        map[string]string `yaml:",omitempty"`

	EvaluateEvery    string `yaml:"evaluate_every"`
	For              string
	OnNoData         string `yaml:"on_no_data"`
	OnExecutionError string `yaml:"on_execution_error"`

	If      []AlertCondition
	Targets []AlertTarget
}

func (a Alert) toOptions() ([]alert.Option, error) { _ = "STUB: not implemented"; return nil, nil }

func (a Alert) targetOptions() ([]alert.Option, error) { _ = "STUB: not implemented"; return nil, nil }

func (a Alert) noDataOption() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

func (a Alert) executionErrorOption() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

type AlertCondition struct {
	Operand *string `yaml:"operand,omitempty"`

	// Query reducers, only one should be used
	Avg         *string `yaml:"avg,omitempty"`
	Sum         *string `yaml:"sum,omitempty"`
	Count       *string `yaml:"count,omitempty"`
	Last        *string `yaml:"last,omitempty"`
	Min         *string `yaml:"min,omitempty"`
	Max         *string `yaml:"max,omitempty"`
	Median      *string `yaml:"median,omitempty"`
	Diff        *string `yaml:"diff,omitempty"`
	PercentDiff *string `yaml:"percent_diff,omitempty"`

	HasNoValue   bool       `yaml:"has_no_value,omitempty"`
	Above        *float64   `yaml:",omitempty"`
	Below        *float64   `yaml:",omitempty"`
	OutsideRange [2]float64 `yaml:"outside_range,omitempty,flow"`
	WithinRange  [2]float64 `yaml:"within_range,omitempty,flow"`
}

func (c AlertCondition) toOption() (alert.Option, error) {
	_ = "STUB: not implemented"
	return *new(alert.Option), nil
}

func (c AlertCondition) queryReducer() (alert.QueryReducer, string, error) {
	_ = "STUB: not implemented"
	return *new(alert.QueryReducer), "", nil
}

func (c AlertCondition) toThresholdOption() (alert.ConditionEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(alert.ConditionEvaluator), nil
}
