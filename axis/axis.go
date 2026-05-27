package axis

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure an axis.
type Option func(axis *Axis)

// Axis represents a visualization axis.
type Axis struct {
	Builder *sdk.Axis
}

// New creates a new Axis configuration.
func New(options ...Option) *Axis { _ = "STUB: not implemented"; return nil }

// Unit sets the unit of the data displayed on this axis.
func Unit(unit string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide makes the axis hidden.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }

// LogBase allows to change the logarithmic scale used to display the values.
func LogBase(base int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets the label on this axis.
func Label(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Min sets the minimum value expected on this axis.
func Min(min float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Max sets the maximum value expected on this axis.
func Max(max float64) Option { _ = "STUB: not implemented"; return *new(Option) }
