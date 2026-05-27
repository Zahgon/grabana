package text

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a textbox variable.
type Option func(constant *Text)

// Text represents a "textbox" templated variable.
type Text struct {
	Builder sdk.TemplateVar
}

// New creates a new "query" templated variable.
func New(name string, options ...Option) *Text { _ = "STUB: not implemented"; return nil }

// Label sets the label of the variable.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// HideLabel ensures that this variable's label will not be displayed.
func HideLabel() Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide ensures that the variable will not be displayed.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }
