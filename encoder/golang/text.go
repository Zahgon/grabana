package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeText(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func textContent(panel *sdk.TextPanel) string { _ = "STUB: not implemented"; return "" }

func textMode(panel *sdk.TextPanel) string { _ = "STUB: not implemented"; return "" }

func textQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
