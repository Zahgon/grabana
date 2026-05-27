package cloudwatch

import (
	"github.com/K-Phoen/grabana/datasource"
	"github.com/K-Phoen/sdk"
)

var _ datasource.Datasource = CloudWatch{}

type CloudWatch struct {
	builder *sdk.Datasource
}

type Option func(datasource *CloudWatch) error

func New(name string, options ...Option) (CloudWatch, error) {
	_ = "STUB: not implemented"
	return *new(CloudWatch), nil
}

func (datasource CloudWatch) Name() string { _ = "STUB: not implemented"; return "" }

func (datasource CloudWatch) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
