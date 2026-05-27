package links

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a panel link.
type Option func(link *Link)

// Link represents a panel link.
type Link struct {
	Builder sdk.Link
}

// New creates a new logs panel.
func New(title string, url string, options ...Option) Link {
	_ = "STUB: not implemented"
	return *new(Link)
}

// OpenBlank configures the link to open in a new tab.
func OpenBlank() Option { _ = "STUB: not implemented"; return *new(Option) }
