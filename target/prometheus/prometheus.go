package prometheus

// FormatMode switches between Table, Time series, or Heatmap. Table will only work
// in the Table panel. Heatmap is suitable for displaying metrics of the
// Histogram type on a Heatmap panel. Under the hood, it converts cumulative
// histograms to regular ones and sorts series by the bucket bound.
type FormatMode string

const (
	FormatTable      FormatMode = "table"
	FormatHeatmap    FormatMode = "heatmap"
	FormatTimeSeries FormatMode = "time_series"
)

// Option represents an option that can be used to configure a prometheus query.
type Option func(target *Prometheus)

// Prometheus represents a prometheus query.
type Prometheus struct {
	Ref            string
	Hidden         bool
	Expr           string
	IntervalFactor int
	Interval       string
	Step           int
	LegendFormat   string
	Instant        bool
	Format         string
}

// New creates a new prometheus query.
func New(query string, options ...Option) *Prometheus { _ = "STUB: not implemented"; return nil }

// Legend sets the legend format.
func Legend(legend string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Ref sets the reference ID for this query.
func Ref(ref string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Hide the query. Grafana does not send hidden queries to the data source,
// but they can still be referenced in alerts.
func Hide() Option { _ = "STUB: not implemented"; return *new(Option) }

// Instant marks the query as "instant, which means Prometheus will only return the latest scrapped value.
func Instant() Option { _ = "STUB: not implemented"; return *new(Option) }

// Format indicates how the data should be returned.
func Format(format FormatMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// IntervalFactor sets the resolution factor.
func IntervalFactor(factor int) Option { _ = "STUB: not implemented"; return *new(Option) }
