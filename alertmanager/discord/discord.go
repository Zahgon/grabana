package discord

import (
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a "discord"
// contact point type.
type Option func(contactType *discordType)

type discordType struct {
	builder *sdk.ContactPointType
}

// With creates an Opsgenie contact point type with the given settings.
func With(webhookURL string, opts ...Option) alertmanager.ContactPointOption {
	_ = "STUB: not implemented"
	return *new(alertmanager.ContactPointOption)
}

// UseDiscordUsername uses the username configured in Discord's webhook settings.
// Otherwise, the username will be 'Grafana'.
func UseDiscordUsername() Option { _ = "STUB: not implemented"; return *new(Option) }
