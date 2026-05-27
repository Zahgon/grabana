package tempo

import (
	"github.com/K-Phoen/grabana/datasource"
	"github.com/K-Phoen/sdk"
)

var _ datasource.Datasource = Tempo{}

type Tempo struct {
	builder *sdk.Datasource
}

type Option func(datasource *Tempo)
type TraceToLogsOption func(settings map[string]interface{})

func New(name string, url string, options ...Option) Tempo {
	_ = "STUB: not implemented"
	return *new(Tempo)
}

func (datasource Tempo) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource Tempo) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
