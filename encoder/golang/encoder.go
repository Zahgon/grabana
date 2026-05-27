package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
	"go.uber.org/zap"
)

const packageImportPath = "github.com/K-Phoen/grabana"
const sdkImportPath = "github.com/K-Phoen/sdk"

type Encoder struct {
	logger *zap.Logger
}

func NewEncoder(logger *zap.Logger) *Encoder { _ = "STUB: not implemented"; return nil }

func (encoder *Encoder) EncodeDashboard(dashboard sdk.Board) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: links, annotations

func (encoder *Encoder) encodeGeneralSettings(dashboard sdk.Board) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// TODO: timezone

func (encoder *Encoder) encodePanels(dashboard sdk.Board) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

func (encoder *Encoder) encodeDataPanel(panel sdk.Panel) (jen.Code, bool) {
	_ = "STUB: not implemented"
	return *new(jen.Code), false
}

/*
	case "singlestat":
		return encoder.encodeSingleStat(panel), true
	case "table":
		return encoder.encodeTable(panel), true
*/

func dashboardQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
