package decoder

import (
	"fmt"

	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/row"
)

var ErrPanelNotConfigured = fmt.Errorf("panel not configured")
var ErrInvalidTimezone = fmt.Errorf("invalid timezone")

type DashboardModel struct {
	Title           string
	Slug            string `yaml:",omitempty"`
	UID             string `yaml:"uid,omitempty"`
	Editable        bool
	SharedCrosshair bool `yaml:"shared_crosshair"`
	Tags            []string
	AutoRefresh     string `yaml:"auto_refresh"`

	Time     [2]string
	Timezone string `yaml:",omitempty"`

	TagsAnnotation []dashboard.TagAnnotation `yaml:"tags_annotations,omitempty"`
	Variables      []DashboardVariable       `yaml:",omitempty"`
	ExternalLinks  []DashboardExternalLink   `yaml:"external_links,omitempty"`
	DashboardLinks []DashboardInternalLink   `yaml:"dashboard_links,omitempty"`

	Rows []DashboardRow
}

func (d *DashboardModel) ToBuilder() (dashboard.Builder, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Builder), nil
}

func (d *DashboardModel) sharedCrossHair() dashboard.Option {
	_ = "STUB: not implemented"
	return *new(dashboard.Option)
}

func (d *DashboardModel) editable() dashboard.Option {
	_ = "STUB: not implemented"
	return *new(dashboard.Option)
}

type DashboardPanel struct {
	Graph      *DashboardGraph      `yaml:",omitempty"`
	Table      *DashboardTable      `yaml:",omitempty"`
	SingleStat *DashboardSingleStat `yaml:"single_stat,omitempty"`
	Stat       *DashboardStat       `yaml:"stat,omitempty"`
	Text       *DashboardText       `yaml:",omitempty"`
	Heatmap    *DashboardHeatmap    `yaml:",omitempty"`
	TimeSeries *DashboardTimeSeries `yaml:"timeseries,omitempty"`
	Logs       *DashboardLogs       `yaml:"logs,omitempty"`
	Gauge      *DashboardGauge      `yaml:"gauge,omitempty"`
}

func (panel DashboardPanel) toOption() (row.Option, error) {
	_ = "STUB: not implemented"
	return *new(row.Option), nil
}
