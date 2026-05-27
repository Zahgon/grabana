package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeGauge(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// todo: thresholds

func (encoder *Encoder) encodeGaugeSettings(panel sdk.Panel) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

func (encoder *Encoder) encodeGaugeOrientation(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func (encoder *Encoder) encodeGaugeValueType(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func gaugeQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
