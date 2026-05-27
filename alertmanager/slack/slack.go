package slack

import (
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a "slack"
// contact point type.
type Option func(contactType *slackType)

type slackType struct {
	builder *sdk.ContactPointType
}

// Webhook creates a Slack contact point type that sends alerts to a Slack webhook.
// See https://api.slack.com/messaging/webhooks
func Webhook(webhookURL string, opts ...Option) alertmanager.ContactPointOption {
	_ = "STUB: not implemented"
	return *new(alertmanager.ContactPointOption)
}

// Title defines a templated title that will be sent in Slack messages.
func Title(templatedTitle string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Body defines the body that will be sent in Slack messages.
func Body(body string) Option { _ = "STUB: not implemented"; return *new(Option) }
