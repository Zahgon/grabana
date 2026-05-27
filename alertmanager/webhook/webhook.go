package webhook

import (
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a "webhook"
// contact point type.
type Option func(contactType *contactType)

type contactType struct {
	builder *sdk.ContactPointType
}

// Call creates a "webhook" contact point type.
func Call(url string, opts ...Option) alertmanager.ContactPointOption {
	_ = "STUB: not implemented"
	return *new(alertmanager.ContactPointOption)
}

// Method defines the HTTP method used to call the webhook. Should be POST or PUT.
func Method(method string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Credentials sets the credentials used to call the webhook.
func Credentials(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// MaxAlerts sets the maximum number of alerts to include in a single call.
// Remaining alerts in the same batch will be ignored above this number.
// 0 means no limit.
func MaxAlerts(max int) Option { _ = "STUB: not implemented"; return *new(Option) }
