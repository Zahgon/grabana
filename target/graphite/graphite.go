package graphite

import "github.com/K-Phoen/sdk"

// Option represents an option that can be used to configure a graphite query.
type Option func(target *Graphite)

// Graphite represents a graphite query.
type Graphite struct {
	Builder *sdk.Target
}

// New creates a new Graphite query.
func New(query string, options ...Option) *Graphite { _ = "STUB: not implemented"; return nil }

// Ref sets the reference ID for this query.
func Ref(ref string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide the query. Grafana does not send hidden queries to the data source,
// but they can still be referenced in alerts.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }
