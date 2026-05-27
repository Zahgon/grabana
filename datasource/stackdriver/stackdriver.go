package stackdriver

import (
	"github.com/K-Phoen/grabana/datasource"
	"github.com/K-Phoen/sdk"
)

var _ datasource.Datasource = Stackdriver{}

type Stackdriver struct {
	builder *sdk.Datasource
}

type Option func(datasource *Stackdriver) error

func New(name string, options ...Option) (Stackdriver, error) {
	_ = "STUB: not implemented"
	return *new(Stackdriver), nil
}

func (datasource Stackdriver) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource Stackdriver) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
