package interval

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure an interval.
type Option func(interval *Interval)

// ValuesList represent a list of options for an interval variable.
type ValuesList []string

// Interval represents a "interval" templated variable.
type Interval struct {
	Builder sdk.TemplateVar
}

// New creates a new "interval" templated variable.
func New(name string, options ...Option) *Interval { _ = "STUB: not implemented"; return nil }

// Values sets the possible values for the variable.
func Values(values ValuesList) Option { _ = "STUB: not implemented"; return *new(Option) }

// Default sets the default value of the variable.
func Default(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets the label of the variable.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// HideLabel ensures that this variable's label will not be displayed.
func HideLabel() Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide ensures that the variable will not be displayed.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }
