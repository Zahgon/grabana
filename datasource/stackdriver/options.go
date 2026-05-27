package stackdriver

// Default configures this datasource to be the default one.
func Default() Option { _ = "STUB: not implemented"; return *new(Option) }

// GCEAuthentication uses GCE default Service Account to authenticate to Stackdriver API.
func GCEAuthentication() Option { _ = "STUB: not implemented"; return *new(Option) }

// JWTAuthentication uses the given ServiceAccount key file to authenticate to Stackdriver API.
func JWTAuthentication(jwt string) Option { _ = "STUB: not implemented"; return *new(Option) }
