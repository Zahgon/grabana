package constant

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a constant.
type Option func(constant *Constant)

// ValuesMap represent a "label" to "value" map of options for a constant variable.
type ValuesMap map[string]string

func (values ValuesMap) asQuery() string { _ = "STUB: not implemented"; return "" }

func (values ValuesMap) labelFor(value string) *sdk.StringSliceString {
	_ = "STUB: not implemented"
	return nil
}

// Constant represents a "constant" templated variable.
type Constant struct {
	Builder sdk.TemplateVar
	values  ValuesMap
}

// New creates a new "constant" templated variable.
func New(name string, options ...Option) *Constant { _ = "STUB: not implemented"; return nil }

// Values sets the possible values for the variable.
func Values(values ValuesMap) Option { _ = "STUB: not implemented"; return *new(Option) }

// Default sets the default value of the variable.
func Default(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets the label of the variable.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// HideLabel ensures that this variable's label will not be displayed.
func HideLabel() Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide ensures that the variable will not be displayed.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }
