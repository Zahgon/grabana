package alertmanager

import (
	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure an
// alert manager.
type Option func(manager *Manager)

// Manager represents an alert manager.
type Manager struct {
	builder *sdk.AlertManager
}

// New creates a new alert manager.
func New(opts ...Option) *Manager { _ = "STUB: not implemented"; return nil }

// ContactPoints defines the contact points that can receive alerts.
func ContactPoints(contactPoints ...Contact) Option { _ = "STUB: not implemented"; return *new(Option) }

// we must have a default contact point, so we use the first contact point
// if none is already set.

// DefaultContactPoint sets the default contact point to be used when no
// specific routing policy applies.
func DefaultContactPoint(contactPoint string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// DefaultGroupBys sets the default labels that alerts should be grouped by.
func DefaultGroupBys(labels ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Templates defines templates that can be used when sending messages to
// contact points.
// See https://prometheus.io/blog/2016/03/03/custom-alertmanager-templates/
func Templates(templates map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Routing configures the routing policies to apply on alerts.
func Routing(policies ...RoutingPolicy) Option { _ = "STUB: not implemented"; return *new(Option) }

// MarshalJSON implements the encoding/json.Marshaler interface.
func (manager *Manager) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalIndentJSON renders the manager as indented JSON.
func (manager *Manager) MarshalIndentJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
