package influxdb

import (
	"github.com/K-Phoen/sdk"
)

type InfluxQL struct {
	builder *sdk.Datasource
}

type Option func(datasource *InfluxQL) error

func New(name, url string, options ...Option) (InfluxQL, error) {
	_ = "STUB: not implemented"
	return *new(InfluxQL), nil
}

func (datasource InfluxQL) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource InfluxQL) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
