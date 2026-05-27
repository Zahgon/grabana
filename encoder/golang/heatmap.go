package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeHeatmap(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// DataFormat

// ShowZeroBuckets/HideZeroBuckets

// HighlightCards/NoHighlightCards

// ReverseYBuckets

// HideXAxis

// Legend()

func (encoder *Encoder) encodeHeatmapYAxis(panel *sdk.HeatmapPanel) jen.Code {
	_ = "STUB: not implemented"
	return *

	// Unit
	new(jen.Code)
}

// Decimals

// Min

// Max

func (encoder *Encoder) encodeHeatmapTooltip(panel *sdk.HeatmapPanel) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// HideTooltip

// HideTooltipHistogram

func heatmapAxisQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }

func heatmapQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
