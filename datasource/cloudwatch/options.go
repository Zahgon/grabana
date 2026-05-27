package cloudwatch

// DefaultAuth relies on the AWS SDK default authentication to authenticate to the CloudWatch service.
func DefaultAuth() Option { _ = "STUB: not implemented"; return *new(Option) }

// AccessSecretAuth relies on an access and secret key to authenticate to the CloudWatch service.
func AccessSecretAuth(accessKey string, secretKey string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Default configures this datasource to be the default one.
func Default() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultRegion sets the default region to use.
// Example: eu-north-1.
func DefaultRegion(region string) Option { _ = "STUB: not implemented"; return *new(Option) }

// AssumeRoleARN specifies the ARN of a role to assume.
// Format: arn:aws:iam:*
func AssumeRoleARN(roleARN string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ExternalID specifies the external identifier of a role to assume in another account.
func ExternalID(externalID string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Endpoint specifies a custom endpoint for the CloudWatch service.
func Endpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

// CustomMetricsNamespaces specifies a list of namespaces for custom metrics.
func CustomMetricsNamespaces(namespaces ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
