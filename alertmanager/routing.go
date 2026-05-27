package alertmanager

import (
	"github.com/K-Phoen/sdk"
)

// RoutingPolicyOption represents an option that can be used to configure a
// routing policy.
type RoutingPolicyOption func(policy *RoutingPolicy)

// RoutingPolicy represents a routing policy.
type RoutingPolicy struct {
	builder *sdk.NotificationRoutingPolicy
}

// Policy defines a routing policy that applies to the given contact point.
// All the options given on this policy will be combined using a logical "AND".
func Policy(contactPoint string, opts ...RoutingPolicyOption) RoutingPolicy {
	_ = "STUB: not implemented"
	return *new(RoutingPolicy)
}

// TagEq defines an equality ("=") constraint between the given tag and value.
func TagEq(tag string, value string) RoutingPolicyOption {
	_ = "STUB: not implemented"
	return *new(RoutingPolicyOption)
}

// TagNeq defines a non-equality ("!=") constraint between the given tag and value.
func TagNeq(tag string, value string) RoutingPolicyOption {
	_ = "STUB: not implemented"
	return *new(RoutingPolicyOption)
}

// TagMatches defines a similarity ("=~") constraint between the given tag and regex.
func TagMatches(tag string, regex string) RoutingPolicyOption {
	_ = "STUB: not implemented"
	return *new(RoutingPolicyOption)
}

// TagNotMatches defines a non-similarity ("!~") constraint between the given tag and regex.
func TagNotMatches(tag string, regex string) RoutingPolicyOption {
	_ = "STUB: not implemented"
	return *new(RoutingPolicyOption)
}
