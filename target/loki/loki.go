package loki

// Option represents an option that can be used to configure a loki query.
type Option func(target *Loki)

// Loki represents a loki query.
type Loki struct {
	Ref          string
	Hidden       bool
	Expr         string
	LegendFormat string
}

// New creates a new prometheus query.
func New(query string, options ...Option) *Loki { _ = "STUB: not implemented"; return nil }

// Legend sets the legend format.
func Legend(legend string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Ref sets the reference ID for this query.
func Ref(ref string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide the query. Grafana does not send hidden queries to the data source,
// but they can still be referenced in alerts.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }
