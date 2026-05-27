package grabana

import (
	"context"
	"io"
	"net/http"
)

// Option represents an option that can be used to configure a client.
type Option func(client *Client)

type requestModifier func(request *http.Request)

// Client represents a Grafana HTTP client.
type Client struct {
	http             *http.Client
	host             string
	requestModifiers []requestModifier
}

// NewClient creates a new Grafana HTTP client, using an API token.
func NewClient(http *http.Client, host string, options ...Option) *Client {
	_ = "STUB: not implemented"
	return nil
}

// WithAPIToken sets up the client to use the given token to authenticate.
func WithAPIToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBasicAuth sets up the client to use the given credentials to authenticate.
func WithBasicAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func (client *Client) modifyRequest(request *http.Request) { _ = "STUB: not implemented"; return }

func (client Client) httpError(resp *http.Response) error { _ = "STUB: not implemented"; return nil }

func (client Client) delete(ctx context.Context, path string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client Client) sendJSON(ctx context.Context, method string, path string, body []byte) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client Client) get(ctx context.Context, path string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client Client) url(path string) string { _ = "STUB: not implemented"; return "" }

func decodeJSON(input io.Reader, data interface{}) error { _ = "STUB: not implemented"; return nil }
