package axis

import (
	"github.com/K-Phoen/sdk"
)

// PlacementMode represents the axis display placement mode.
type PlacementMode string

const (
	Hidden PlacementMode = "hidden"
	Auto   PlacementMode = "auto"
	Left   PlacementMode = "left"
	Right  PlacementMode = "right"
)

// ScaleMode represents the axis scale distribution.
type ScaleMode uint8

const (
	Linear ScaleMode = iota
	Log2
	Log10
)

// Option represents an option that can be used to configure an axis.
type Option func(axis *Axis) error

// Axis represents a visualization axis.
type Axis struct {
	fieldConfig *sdk.FieldConfig
}

// New creates a new Axis configuration.
func New(fieldConfig *sdk.FieldConfig, options ...Option) (*Axis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Placement defines how the axis should be placed in the panel.
func Placement(placement PlacementMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// SoftMin defines a soft minimum value for the axis.
func SoftMin(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// SoftMax defines a soft maximum value for the axis.
func SoftMax(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Min defines a hard minimum value for the axis.
func Min(value float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Max defines a hard maximum value for the axis.
func Max(value float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Unit sets the unit of the data displayed in this series.
func Unit(unit string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Scale sets the scale to use for the Y-axis values..
func Scale(mode ScaleMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets a Y-axis text label.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Decimals sets how many decimal points should be displayed.
func Decimals(decimals int) Option { _ = "STUB: not implemented"; return *new(Option) }
