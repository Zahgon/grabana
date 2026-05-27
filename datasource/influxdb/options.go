package influxdb

import (
	"time"
)

type Access string

const (
	Proxy   Access = "proxy"
	Browser Access = "direct"
)

// Default configures this datasource to be the default one.
func Default() Option { _ = "STUB: not implemented"; return *new(Option) }

// HTTPMethod defines the Method used to query the database (GET or POST HTTP verb).
// The POST verb allows heavy queries that would return an error using the GET verb.
// Default is GET.
func HTTPMethod(method string) Option { _ = "STUB: not implemented"; return *new(Option) }

// AccessMode controls how requests to the data source will be handled. Proxy
// should be the preferred way if nothing else is stated. Browser will let your
// browser send the requests (deprecated).
func AccessMode(mode Access) Option { _ = "STUB: not implemented"; return *new(Option) }

// KeepCookies controls the cookies that will be forwarded to the data source.
// All other cookies will be deleted.
func KeepCookies(cookies []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Timeout sets the timeout for HTTP requests.
func Timeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// Database sets the ID of the bucket you want to query from,
// copied from the Buckets page of the InfluxDB UI.
func Database(database string) Option { _ = "STUB: not implemented"; return *new(Option) }

// User sets username to use to sign into InfluxDB.
func User(user string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Password sets token you use to query the selected bucked,
// copied from the Tokens page of the InfluxDB UI.
func Password(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// MinTimeInterval defines a lower limit for the auto group by time interval.
// Recommended to be set to write frequency, for example 1m if your data is written every minute.
func MinTimeInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxSeries limits the number of series/tables that Grafana processes.
// Lower this number to prevent abuse, and increase it if you have lots of small time series
// and not all are shown. Defaults to 1000.
func MaxSeries(max int) Option { _ = "STUB: not implemented"; return *new(Option) }

// BasicAuth configures basic authentication for this datasource.
func BasicAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCredentials joins credentials such as cookies or auth headers to cross-site requests.
func WithCredentials() Option { _ = "STUB: not implemented"; return *new(Option) }

// SkipTLSVerify disables verification of SSL certificates.
func SkipTLSVerify() Option { _ = "STUB: not implemented"; return *new(Option) }

// ForwardOauthIdentity forward the user's upstream OAuth identity to the datasource.
func ForwardOauthIdentity() Option { _ = "STUB: not implemented"; return *new(Option) }

// TLSClientAuth enables TLS client side authentication. Expects PEM encoded content.
func TLSClientAuth(cert string, key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCACert allows to provide a PEM encoded CA certificate to trust for this data source.
func WithCACert(cert string) Option { _ = "STUB: not implemented"; return *new(Option) }

func multiOption(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func setJSONData(key string, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func setSecureJSONData(key string, value interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func invalidArgument(err error) Option { _ = "STUB: not implemented"; return *new(Option) }
