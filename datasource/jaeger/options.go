package jaeger

import "time"

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

// WithNodeGraph enables the Node Graph visualization in the trace viewer.
func WithNodeGraph() Option { _ = "STUB: not implemented"; return *new(Option) }

// TraceToLogs defines how to navigate from a trace span to the selected datasource logs.
func TraceToLogs(logsDatasourceUID string, options ...TraceToLogsOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Tags defines tags that will be used in the Loki query.
// Default tags: 'cluster', 'hostname', 'namespace', 'pod'.
func Tags(tags ...string) TraceToLogsOption {
	_ = "STUB: not implemented"
	return *new(TraceToLogsOption)
}

// SpanStartShift shifts the start time of the span.
// Default 0 (Time units can be used here, for example: 5s, 1m, 3h)
func SpanStartShift(shift time.Duration) TraceToLogsOption {
	_ = "STUB: not implemented"
	return *new(TraceToLogsOption)
}

// SpanEndShift shifts the start time of the span.
// Default 0 (Time units can be used here, for example: 5s, 1m, 3h)
func SpanEndShift(shift time.Duration) TraceToLogsOption {
	_ = "STUB: not implemented"
	return *new(TraceToLogsOption)
}

// FilterByTrace filters logs by Trace ID. Appends '|=<trace id>' to the query.
func FilterByTrace() TraceToLogsOption { _ = "STUB: not implemented"; return *new(TraceToLogsOption) }

// FilterBySpan filters logs by Trace ID. Appends '|=<trace id>' to the query.
func FilterBySpan() TraceToLogsOption { _ = "STUB: not implemented"; return *new(TraceToLogsOption) }
