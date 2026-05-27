package email

import (
	"github.com/K-Phoen/grabana/alertmanager"
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure an "email"
// contact point type.
type Option func(contactType *emailType)

type emailType struct {
	builder *sdk.ContactPointType
}

// To creates an "email" contact point type.
func To(emails []string, opts ...Option) alertmanager.ContactPointOption {
	_ = "STUB: not implemented"
	return *new(alertmanager.ContactPointOption)
}

// Single send a single email to all recipients.
func Single() Option { _ = "STUB: not implemented"; return *new(Option) }

// Message sets an optional message that will be included in the email.
// Variables are allowed.
func Message(content string) Option { _ = "STUB: not implemented"; return *new(Option) }
