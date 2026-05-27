package golang

import (
	"github.com/K-Phoen/jennifer/jen"
)

type RowIR struct {
	Title     string
	RepeatFor *string
	Collapsed bool
	Panels    []jen.Code
}

func (encoder *Encoder) encodeRow(row RowIR) *jen.Statement { _ = "STUB: not implemented"; return nil }

func rowQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
