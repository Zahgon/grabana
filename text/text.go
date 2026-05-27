package text

import (
	"github.com/K-Phoen/grabana/links"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a text panel.
type Option func(text *Text) error

// Text represents a text panel.
type Text struct {
	Builder *sdk.Panel
}

// New creates a new text panel.
func New(title string, options ...Option) (*Text, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Links adds links to be displayed on this panel.
func Links(panelLinks ...links.Link) Option { _ = "STUB: not implemented"; return *new(Option) }

// HTML sets the content of the panel, to be rendered as HTML.
func HTML(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Markdown sets the content of the panel, to be rendered as markdown.
func Markdown(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Span sets the width of the panel, in grid units. Should be a positive
// number between 1 and 12. Example: 6.
func Span(span float32) Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the panel, in pixels. Example: "400px".
func Height(height string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description annotates the current visualization with a human-readable description.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Transparent makes the background transparent.
func Transparent() Option { _ = "STUB: not implemented"; return *new(Option) }
