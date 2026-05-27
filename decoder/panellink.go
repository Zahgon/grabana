package decoder

import (
	"github.com/K-Phoen/grabana/links"
)

type DashboardPanelLinks []DashboardPanelLink

func (collection DashboardPanelLinks) toModel() []links.Link { _ = "STUB: not implemented"; return nil }

type DashboardPanelLink struct {
	Title        string
	URL          string `yaml:"url"`
	OpenInNewTab bool   `yaml:"open_in_new_tab,omitempty"`
}

func (l DashboardPanelLink) toModel() links.Link {
	_ = "STUB: not implemented"
	return *new(links.Link)
}
