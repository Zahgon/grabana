package dashboard

import (
	// We're not using it for security stuff, so it's fine.
	//nolint:gosec

	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/variable/constant"
	"github.com/K-Phoen/grabana/variable/custom"
	"github.com/K-Phoen/grabana/variable/datasource"
	"github.com/K-Phoen/grabana/variable/interval"
	"github.com/K-Phoen/grabana/variable/query"
	"github.com/K-Phoen/grabana/variable/text"
	"github.com/K-Phoen/sdk"
)

// TagAnnotation describes an annotation represented as a Tag.
// See https://grafana.com/docs/grafana/latest/reference/annotations/#query-by-tag
type TagAnnotation struct {
	Name       string
	Datasource string
	IconColor  string   `yaml:"color"`
	Tags       []string `yaml:",flow"`
}

// Option represents an option that can be used to configure a
// dashboard.
type Option func(dashboard *Builder) error

// TimezoneOption represents a possible value for the dashboard's timezone
// configuration.
type TimezoneOption string

// DefaultTimezone sets the dashboard's timezone to the default one used by
// Grafana.
const DefaultTimezone TimezoneOption = ""

// UTC sets the dashboard's timezone to UTC.
const UTC TimezoneOption = "utc"

// Browser sets the dashboard's timezone to the browser's one.
const Browser TimezoneOption = "browser"

// Builder is the main builder used to configure dashboards.
type Builder struct {
	board  *sdk.Board
	alerts []*alert.Alert
}

// New creates a new dashboard builder.
func New(title string, options ...Option) (Builder, error) {
	_ = "STUB: not implemented"
	return *new(Builder), nil
}

func defaults() []Option { _ = "STUB: not implemented"; return nil }

func defaultTimePicker() Option { _ = "STUB: not implemented"; return *new(Option) }

// MarshalJSON implements the encoding/json.Marshaler interface.
//
// This method can be used to render the dashboard as JSON
// which your configuration management tool of choice can then feed into
// Grafana's dashboard via its provisioning support.
// See https://grafana.com/docs/grafana/latest/administration/provisioning/#dashboards
func (builder *Builder) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalIndentJSON renders the dashboard as indented JSON
// which your configuration management tool of choice can then feed into
// Grafana's dashboard via its provisioning support.
// See https://grafana.com/docs/grafana/latest/administration/provisioning/#dashboards
func (builder *Builder) MarshalIndentJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Alerts returns all the alerts defined in this dashboard.
func (builder *Builder) Alerts() []*alert.Alert { _ = "STUB: not implemented"; return nil }

// Internal.
func (builder *Builder) Internal() *sdk.Board { _ = "STUB: not implemented"; return nil }

// VariableAsConst adds a templated variable, defined as a set of constant
// values.
// See https://grafana.com/docs/grafana/latest/reference/templating/#variable-types
func VariableAsConst(name string, options ...constant.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ID sets the ID used by the dashboard.
func ID(id uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// UID sets the UID used by the dashboard.
func UID(uid string) Option { _ = "STUB: not implemented"; return *new(Option) }

// We're not using it for security stuff, so it's fine.
//nolint:gosec

// Slug sets the Slug used by the dashboard.
func Slug(slug string) Option { _ = "STUB: not implemented"; return *new(Option) }

// VariableAsCustom adds a templated variable, defined as a set of custom
// values.
// See https://grafana.com/docs/grafana/latest/reference/templating/#variable-types
func VariableAsCustom(name string, options ...custom.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// VariableAsInterval adds a templated variable, defined as an interval.
// See https://grafana.com/docs/grafana/latest/reference/templating/#variable-types
func VariableAsInterval(name string, options ...interval.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// VariableAsQuery adds a templated variable, defined as a query.
// See https://grafana.com/docs/grafana/latest/reference/templating/#variable-types
func VariableAsQuery(name string, options ...query.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// VariableAsDatasource adds a templated variable, defined as a datasource.
// See https://grafana.com/docs/grafana/latest/variables/variable-types/add-data-source-variable/
func VariableAsDatasource(name string, options ...datasource.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// VariableAsText adds a templated variable, defined as a free text input.
// See https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#add-a-text-box-variable
func VariableAsText(name string, options ...text.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ExternalLinks adds a dashboard-level external links.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/manage-dashboard-links/#add-a-url-link-to-a-dashboard
func ExternalLinks(links ...ExternalLink) Option { _ = "STUB: not implemented"; return *new(Option) }

// DashboardLinks adds a dashboard-level links to other dashboards.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/manage-dashboard-links/#dashboard-links
func DashboardLinks(links ...DashboardLink) Option {
	_ = "STUB: not implemented" //nolint:revive
	return *new(Option)
}

// Row adds a row to the dashboard.
func Row(title string, options ...row.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// TagsAnnotation adds a new source of annotation for the dashboard.
func TagsAnnotation(annotation TagAnnotation) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Editable marks the dashboard as editable.
func Editable() Option { _ = "STUB: not implemented"; return *new(Option) }

// ReadOnly marks the dashboard as non-editable.
func ReadOnly() Option { _ = "STUB: not implemented"; return *new(Option) }

// SharedCrossHair configures the graph tooltip to be shared across panels.
func SharedCrossHair() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultTooltip configures the graph tooltip NOT to be shared across panels.
func DefaultTooltip() Option { _ = "STUB: not implemented"; return *new(Option) }

// Tags adds the given set of tags to the dashboard.
func Tags(tags []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// AutoRefresh defines the auto-refresh interval for the dashboard.
func AutoRefresh(interval string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Time defines the default time range for the dashboard, e.g. from "now-6h" to
// "now".
func Time(from, to string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Timezone defines the default timezone for the dashboard, e.g. "utc".
func Timezone(timezone TimezoneOption) Option { _ = "STUB: not implemented"; return *new(Option) }
