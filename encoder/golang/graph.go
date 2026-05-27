package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeGraph(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// TODO: SeriesOverride()

// Null

// LineWidth

// Fill

// PointRadius

// Staircase

func (encoder *Encoder) encodeGraphAxis(graphOptName string, axis sdk.Axis) jen.Code {
	_ = "STUB: not implemented"
	return *

	// Unit
	new(jen.Code)
}

// Hide

// LogBase

// Label

// Min

// Max

func (encoder *Encoder) encodeGraphLegend(legend sdk.Legend) jen.Code {
	_ = "STUB: not implemented"
	return *

	// Hidden legend?
	new(jen.Code)
}

func (encoder *Encoder) encodeGraphDraw(panel sdk.GraphPanel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func graphQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
