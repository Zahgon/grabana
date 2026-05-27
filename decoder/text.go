package decoder

import (
	"github.com/K-Phoen/grabana/row"
)

type DashboardText struct {
	Title       string
	Description string              `yaml:",omitempty"`
	Span        float32             `yaml:",omitempty"`
	Height      string              `yaml:",omitempty"`
	Transparent bool                `yaml:",omitempty"`
	Links       DashboardPanelLinks `yaml:",omitempty"`
	HTML        string              `yaml:",omitempty"`
	Markdown    string              `yaml:",omitempty"`
}

func (textPanel DashboardText) toOption() row.Option {
	_ = "STUB: not implemented"
	return *new(row.Option)
}
