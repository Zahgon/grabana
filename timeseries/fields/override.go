package fields

import (
	"github.com/K-Phoen/grabana/timeseries/axis"
	"github.com/K-Phoen/sdk"
)

// StackMode configures mode of series stacking.
// FIXME: copied here to avoid circular imports with parent package
type StackMode string

const (
	// Unstacked will not stack series
	Unstacked StackMode = "none"
	// NormalStack will stack series as absolute numbers
	NormalStack StackMode = "normal"
	// PercentStack will stack series as percents
	PercentStack StackMode = "percent"
)

type OverrideOption func(field *sdk.FieldConfigOverride)

// Unit overrides the unit.
func Unit(unit string) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

// FillOpacity overrides the opacity.
func FillOpacity(opacity int) OverrideOption {
	_ = "STUB: not implemented"
	return *new(OverrideOption)
}

// FixedColorScheme forces the use of a fixed color scheme.
func FixedColorScheme(color string) OverrideOption {
	_ = "STUB: not implemented"
	return *new(OverrideOption)
}

// NegativeY flips the results to negative values on the Y axis.
func NegativeY() OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

// AxisPlacement overrides how the axis should be placed in the panel.
func AxisPlacement(placement axis.PlacementMode) OverrideOption {
	_ = "STUB: not implemented"
	return *new(OverrideOption)
}

// Stack overrides if the series should be stacked and using which mode (default not stacked).
func Stack(mode StackMode) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }
