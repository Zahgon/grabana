package custom

import (
	"github.com/K-Phoen/sdk"
)

const All = "$__all"

// Option represents an option that can be used to configure a custom variable.
type Option func(constant *Custom)

// ValuesMap represent a "label" to "value" map of options for a custom variable.
type ValuesMap map[string]string

func (values ValuesMap) asQuery() string { _ = "STUB: not implemented"; return "" }

func (values ValuesMap) labelFor(value string) *sdk.StringSliceString {
	_ = "STUB: not implemented"
	return nil
}

// Custom represents a "custom" templated variable.
type Custom struct {
	Builder sdk.TemplateVar
	values  ValuesMap
}

// New creates a new "custom" templated variable.
func New(name string, options ...Option) *Custom { _ = "STUB: not implemented"; return nil }

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

// Multiple allows several values to be selected.
func Multiple() Option { _ = "STUB: not implemented"; return *new(Option) }

// IncludeAll adds an option to allow all values to be selected.
func IncludeAll() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultAll selects "All" values by default.
func DefaultAll() Option { _ = "STUB: not implemented"; return *new(Option) }

// AllValue define the value used when selecting the "All" option.
func AllValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }
