package jaeger

import (
	"github.com/K-Phoen/grabana/datasource"
	"github.com/K-Phoen/sdk"
)

var _ datasource.Datasource = Jaeger{}

type Jaeger struct {
	builder *sdk.Datasource
}

type Option func(datasource *Jaeger)
type TraceToLogsOption func(settings map[string]interface{})

func New(name string, url string, options ...Option) Jaeger {
	_ = "STUB: not implemented"
	return *new(Jaeger)
}

func (datasource Jaeger) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource Jaeger) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
