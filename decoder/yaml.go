package decoder

import (
	"io"

	"github.com/K-Phoen/grabana/dashboard"
)

func UnmarshalYAML(input io.Reader) (dashboard.Builder, error) {
	_ = "STUB: not implemented"
	return *new(dashboard.Builder), nil
}
