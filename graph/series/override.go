package series

import (
	"github.com/K-Phoen/sdk"
)

// OverrideOption represents an option that can be used alter a graph panel series.
type OverrideOption func(series *sdk.SeriesOverride) error

// Alias defines an alias/regex used to identify the series to override.
func Alias(alias string) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

// Color overrides the color for the matched series.
func Color(color string) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

// Dashes enables/disables display of the series using dashes instead of lines.
func Dashes(enabled bool) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

// Lines enables/disables display of the series using dashes instead of dashes.
func Lines(enabled bool) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

func Fill(opacity int) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }

func LineWidth(width int) OverrideOption { _ = "STUB: not implemented"; return *new(OverrideOption) }
