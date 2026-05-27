package grabana

import (
	"context"
	"errors"
)

// ErrAPIKeyNotFound is returned when the given API key can not be found.
var ErrAPIKeyNotFound = errors.New("API key not found")

// APIKeyRole represents a role given to an API key.
type APIKeyRole uint8

const (
	AdminRole APIKeyRole = iota
	EditorRole
	ViewerRole
)

func (role APIKeyRole) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateAPIKeyRequest represents a request made to the API key creation endpoint.
type CreateAPIKeyRequest struct {
	Name          string     `json:"name"`
	Role          APIKeyRole `json:"role"`
	SecondsToLive int        `json:"secondsToLive"`
}

// APIKey represents an API key.
type APIKey struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CreateAPIKey creates a new API key.
func (client *Client) CreateAPIKey(ctx context.Context, request CreateAPIKeyRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeleteAPIKeyByName deletes an API key given its name.
func (client *Client) DeleteAPIKeyByName(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// APIKeys lists active API keys.
func (client *Client) APIKeys(ctx context.Context) (map[string]APIKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
