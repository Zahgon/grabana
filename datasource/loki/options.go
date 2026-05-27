package loki

import "time"

type DerivedField struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	// Used to parse and capture some part of the log message. You can use the captured groups in the template.
	Regex string `json:"matcherRegex"`
	// Used to override the button label when this derived field is found in a log.
	// Optional.
	URLDisplayLabel string `json:"urlDisplayLabel,omitempty"`
	// For internal links
	// Optional.
	DatasourceUID string `json:"datasourceUid,omitempty"`
}

// Default configures this datasource to be the default one.
func Default() Option { _ = "STUB: not implemented"; return *new(Option) }

// Timeout sets the timeout for HTTP requests.
func Timeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// BasicAuth configures basic authentication for this datasource.
func BasicAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

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

// MaximumLines sets the maximum number of lines returned by Loki (default: 1000).
// Increase this value to have a bigger result set for ad-hoc analysis.
// Decrease this limit if your browser becomes sluggish when displaying the
// log results.
func MaximumLines(max int) Option { _ = "STUB: not implemented"; return *new(Option) }

// DerivedFields defines fields can be used to extract new fields from a log
// message and create a link from its value.
func DerivedFields(fields ...DerivedField) Option { _ = "STUB: not implemented"; return *new(Option) }
