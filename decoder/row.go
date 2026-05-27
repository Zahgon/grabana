package decoder

import (
	"github.com/K-Phoen/grabana/dashboard"
)

// DashboardRow represents a dashboard row.
type DashboardRow struct {
	Name      string
	Repeat    string `yaml:"repeat_for,omitempty"`
	Collapse  bool   `yaml:",omitempty"`
	HideTitle bool   `yaml:"hide_title,omitempty"`
	Panels    []DashboardPanel
}

func (r DashboardRow) toOption() (dashboard.Option, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Option), nil
}
