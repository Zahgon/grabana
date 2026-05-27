package axis

// Option represents an option that can be used to configure a Y axis.
type Option func(axis *YAxis)

type SDKAxis struct {
	Decimals    *int     `json:"decimals"`
	Format      string   `json:"format"`
	LogBase     int      `json:"logBase"`
	Show        bool     `json:"show"`
	Max         *string  `json:"max"`
	Min         *string  `json:"min"`
	SplitFactor *float64 `json:"splitFactor"`
}

// YAxis represents the Y axis of a heatmap.
type YAxis struct {
	Builder *SDKAxis
}

// New creates a new YAxis configuration.
func New(options ...Option) *YAxis { _ = "STUB: not implemented"; return nil }

// Unit sets the unit of the data displayed on this axis.
func Unit(unit string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Decimals set the number of decimals to be displayed on the axis.
func Decimals(decimals int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Min sets the minimum value expected on this axis.
func Min(min float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// Max sets the maximum value expected on this axis.
func Max(max float64) Option { _ = "STUB: not implemented"; return *new(Option) }
