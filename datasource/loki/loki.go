package loki

import (
	"github.com/K-Phoen/grabana/datasource"
	"github.com/K-Phoen/sdk"
)

var _ datasource.Datasource = Loki{}

type Loki struct {
	builder *sdk.Datasource
}

type Option func(datasource *Loki)

func New(name string, url string, options ...Option) Loki {
	_ = "STUB: not implemented"
	return *new(Loki)
}

func (datasource Loki) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource Loki) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
