package prometheus

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

// BasicAuth configures basic authentication for this datasource.
func BasicAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AccessMode controls how requests to the data source will be handled. Proxy
// should be the preferred way if nothing else is stated. Browser will let your
// browser send the requests (deprecated).
func AccessMode(mode Access) Option { _ = "STUB: not implemented"; return *new(Option) }

// HTTPMethod sets the method used to query Prometheus. POST is the recommended
// method as it allows bigger queries. Change this to GET if you have a
// Prometheus version older than 2.1 or if POST requests are restricted in your
// network.
func HTTPMethod(method string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ScrapeInterval configures the scrape and evaluation interval. Should be set
// to the typical value in Prometheus (defaults to 15s).
func ScrapeInterval(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// QueryTimeout sets the timeout for queries. Defaults to 60s
func QueryTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// SkipTLSVerify disables verification of SSL certificates.
func SkipTLSVerify() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCertificate sets a self-signed certificate that can be verified against.
func WithCertificate(certificate string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCredentials joins credentials such as cookies or auth headers to cross-site requests.
func WithCredentials() Option { _ = "STUB: not implemented"; return *new(Option) }

// ForwardOauthIdentity forward the user's upstream OAuth identity to the data
// source (Their access token gets passed along).
func ForwardOauthIdentity() Option { _ = "STUB: not implemented"; return *new(Option) }

// ForwardCookies configures a list of cookies that should be forwarded to the
// datasource.
func ForwardCookies(cookies ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Exemplars configures a list of exemplars on this datasource.
func Exemplars(exemplars ...Exemplar) Option { _ = "STUB: not implemented"; return *new(Option) }
