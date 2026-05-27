package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeStat(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// TODO: ColorScheme

func (encoder *Encoder) encodeStatOptions(options sdk.StatOptions) []jen.Code {
	_ = "STUB: not implemented"
	return nil

	// SparkLine
}

// ValueFontSize

// TitleFontSize

// Text

// Orientation

// ValueType

// Automatic calculations

func (encoder *Encoder) encodeStatFieldConfigDefaults(defaults sdk.FieldConfigDefaults) []jen.Code {
	_ = "STUB: not implemented"
	return nil

	// unit
}

// decimals

// sparkline Y min

// sparkline Y max

// NoValue

// TODO: thresholds
// RelativeThresholds/AbsoluteThresholds
/*
	if defaults.Thresholds.Mode != "" && len(defaults.Thresholds.Steps) != 0 {
	}
*/

func statQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
