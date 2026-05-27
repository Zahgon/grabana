package scheme

import (
	"github.com/K-Phoen/sdk"
)

type ColorMode string

const (
	Last ColorMode = "last"
	Min  ColorMode = "min"
	Max  ColorMode = "max"
)

// Option represents an option that can be used to configure an axis.
type Option func(scheme *Scheme)

type Step struct {
	Color string
	Value int
}

// Scheme represents a color scheme.
type Scheme struct {
	fieldConfig *sdk.FieldConfig
}

// New creates a new Scheme configuration.
func New(fieldConfig *sdk.FieldConfig, options ...Option) *Scheme {
	_ = "STUB: not implemented"
	return nil
}

// SingleColor defines the color scheme with a single color.
func SingleColor(color string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ClassicPalette uses the classic palette color scheme.
func ClassicPalette() Option { _ = "STUB: not implemented"; return *new(Option) }

// ThresholdsValue uses the thresholds colors.
func ThresholdsValue(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// GreenYellowRed uses the green-yellow-red color scheme.
func GreenYellowRed(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// YellowRed uses the yellow-red color scheme.
func YellowRed(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// YellowBlue uses the yellow-blue color scheme.
func YellowBlue(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// RedYellowGreen uses the red-yellow-green color scheme.
func RedYellowGreen(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// BlueYellowRed uses the blue-yellow-red color scheme.
func BlueYellowRed(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// BluePurple uses the blue-purple color scheme.
func BluePurple(colorBy ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }
