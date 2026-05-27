package opsgenie

import (
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure an "opsgenie"
// contact point type.
type Option func(contactType *opsgenieType)

// TagForwardMode describes how alert tags should be forwarded to Opsgenie.
type TagForwardMode string

const (
	Tags                   TagForwardMode = "tags"
	ExtraProperties        TagForwardMode = "details"
	TagsAndExtraProperties TagForwardMode = "both"
)

type opsgenieType struct {
	builder *sdk.ContactPointType
}

// With creates an Opsgenie contact point type with the given settings.
func With(apiURL string, apiKey string, opts ...Option) alertmanager.ContactPointOption {
	_ = "STUB: not implemented"
	return *new(alertmanager.ContactPointOption)
}

// AutoClose automatically closes an alert in Opsgenie once it goes back to OK in Grafana.
func AutoClose() Option { _ = "STUB: not implemented"; return *new(Option) }

// OverridePriority allows the alert priority to be set in Opsgenie based on
// the content of the `og_priority` annotation.
func OverridePriority() Option { _ = "STUB: not implemented"; return *new(Option) }

// SentTagsAs defines how alert tags should be forwarded to Opsgenie.
func SentTagsAs(mode TagForwardMode) Option { _ = "STUB: not implemented"; return *new(Option) }
