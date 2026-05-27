package decoder

import (
	"github.com/K-Phoen/grabana/dashboard"
)

type DashboardInternalLink struct {
	Title                 string   `yaml:"title"`
	Tags                  []string `yaml:"tags"`
	AsDropdown            bool     `yaml:"as_dropdown,omitempty"`
	IncludeTimeRange      bool     `yaml:"include_time_range,omitempty"`
	IncludeVariableValues bool     `yaml:"include_variable_values,omitempty"`
	OpenInNewTab          bool     `yaml:"open_in_new_tab,omitempty"`
}

func (l DashboardInternalLink) toModel() dashboard.DashboardLink {
	_ = "STUB: not implemented"
	return *new(dashboard.DashboardLink)
}
