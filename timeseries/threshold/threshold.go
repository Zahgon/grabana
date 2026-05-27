package threshold

import (
	"github.com/K-Phoen/sdk"
)

// Mode represents the type of value used as threshold.
type Mode string

const (
	Percentage Mode = "percentage"
	Absolute   Mode = "absolute"
)

// DisplayStyle represents how the threshold should be visualized.
type DisplayStyle string

const (
	Off             DisplayStyle = "off"
	AsFilledRegions DisplayStyle = "area"
	AsLines         DisplayStyle = "line"
	Both            DisplayStyle = "line+area"
)

// Option represents an option that can be used to configure an axis.
type Option func(threshold *Threshold)

type Step struct {
	Color string
	Value float64
}

// Threshold represents a threshold visualization.
type Threshold struct {
	baseColor   string
	fieldConfig *sdk.FieldConfig
}

// New creates a new Threshold configuration.
func New(fieldConfig *sdk.FieldConfig, options ...Option) *Threshold {
	_ = "STUB: not implemented"
	return nil
}

// Style defines the thresholds display style.
func Style(style DisplayStyle) Option { _ = "STUB: not implemented"; return *new(Option) }

// BaseColor defines the color of the thresholds' base.
func BaseColor(color string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ValueMode defines how to interpret the threshold values.
func ValueMode(mode Mode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Steps defines threshold steps.
func Steps(steps ...Step) Option { _ = "STUB: not implemented"; return *new(Option) }

// Base

// User-defined steps
