package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeTimeseries(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// TODO: overrides

func (encoder *Encoder) encodeTimeseriesAxis(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// label

// decimals

// boundaries

// placement

// scale

func (encoder *Encoder) encodeTimeseriesLegend(legend sdk.TimeseriesLegendOptions) jen.Code {
	_ = "STUB: not implemented"
	return *

	// Hidden legend?
	new(jen.Code)
}

// Display mode

// Placement

// Automatic calculations

func (encoder *Encoder) encodeTimeseriesVizualization(panel sdk.Panel) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// Line interpolation mode

// don't generate code for the default

// Tooltip mode

// don't generate code for the default

// Gradient mode

// don't generate code for the default

// Stacking mode

// don't generate code for the default

func timeseriesQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }

func tsAxisQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
