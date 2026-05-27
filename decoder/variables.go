package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/dashboard"
)

var ErrVariableNotConfigured = fmt.Errorf("variable not configured")
var ErrInvalidHideValue = fmt.Errorf("invalid hide value. Valid values are: 'label', 'variable', empty")

type DashboardVariable struct {
	Interval   *VariableInterval   `yaml:",omitempty"`
	Custom     *VariableCustom     `yaml:",omitempty"`
	Query      *VariableQuery      `yaml:",omitempty"`
	Const      *VariableConst      `yaml:",omitempty"`
	Datasource *VariableDatasource `yaml:",omitempty"`
	Text       *VariableText       `yaml:",omitempty"`
}

func (variable *DashboardVariable) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

type VariableInterval struct {
	Name    string
	Label   string   `yaml:",omitempty"`
	Default string   `yaml:",omitempty"`
	Values  []string `yaml:",flow"`
	Hide    string   `yaml:",omitempty"`
}

func (variable *VariableInterval) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do

type VariableCustom struct {
	Name       string
	Label      string            `yaml:",omitempty"`
	Default    string            `yaml:",omitempty"`
	ValuesMap  map[string]string `yaml:"values_map"`
	IncludeAll bool              `yaml:"include_all"`
	AllValue   string            `yaml:"all_value,omitempty"`
	Hide       string            `yaml:",omitempty"`
	Multiple   bool              `yaml:",omitempty"`
}

func (variable *VariableCustom) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do

type VariableConst struct {
	Name      string
	Label     string            `yaml:",omitempty"`
	Default   string            `yaml:",omitempty"`
	ValuesMap map[string]string `yaml:"values_map"`
	Hide      string            `yaml:",omitempty"`
}

func (variable *VariableConst) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do

type VariableQuery struct {
	Name  string
	Label string `yaml:",omitempty"`

	Datasource string `yaml:",omitempty"`
	Request    string

	Regex      string `yaml:",omitempty"`
	IncludeAll bool   `yaml:"include_all"`
	DefaultAll bool   `yaml:"default_all"`
	AllValue   string `yaml:"all_value,omitempty"`
	Hide       string `yaml:",omitempty"`
	Multiple   bool   `yaml:",omitempty"`
}

func (variable *VariableQuery) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do

type VariableDatasource struct {
	Name  string
	Label string `yaml:",omitempty"`

	Type string

	Regex      string `yaml:",omitempty"`
	IncludeAll bool   `yaml:"include_all"`
	Hide       string `yaml:",omitempty"`
	Multiple   bool   `yaml:",omitempty"`
}

func (variable *VariableDatasource) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do

type VariableText struct {
	Name  string
	Label string `yaml:",omitempty"`
	Hide  string `yaml:",omitempty"`
}

func (variable *VariableText) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}

// Nothing to do
