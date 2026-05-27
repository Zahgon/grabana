package datasource

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a query.
type Option func(constant *Datasource)

const (
	// dashboardLoad will refresh the results every time the dashboard is loaded.
	dashboardLoad int64 = 1
)

// Datasource represents a "datasource" templated variable.
type Datasource struct {
	Builder sdk.TemplateVar
}

// New creates a new "query" templated variable.
func New(name string, options ...Option) *Datasource { _ = "STUB: not implemented"; return nil }

// Type defines the datasource type. Example: "grafana", "stackdriver", "prometheus", ...
func Type(datasourceType string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Regex defines a filter allowing to filter the values returned by the request/query.
func Regex(regex string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets the label of the variable.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// HideLabel ensures that this variable's label will not be displayed.
func HideLabel() Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide ensures that the variable will not be displayed.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }

// Multiple allows several values to be selected.
func Multiple() Option { _ = "STUB: not implemented"; return *new(Option) }

// IncludeAll adds an option to allow all values to be selected.
func IncludeAll() Option { _ = "STUB: not implemented"; return *new(Option) }
