package decoder

import (
	"github.com/K-Phoen/grabana/dashboard"
)

type DashboardExternalLink struct {
	Title                 string
	URL                   string `yaml:"url"`
	Description           string `yaml:",omitempty"`
	Icon                  string `yaml:"icon,omitempty"`
	IncludeTimeRange      bool   `yaml:"include_time_range,omitempty"`
	IncludeVariableValues bool   `yaml:"include_variable_values,omitempty"`
	OpenInNewTab          bool   `yaml:"open_in_new_tab,omitempty"`
}

func (l DashboardExternalLink) toModel() dashboard.ExternalLink {
	_ = "STUB: not implemented"
	return *new(dashboard.ExternalLink)
}
