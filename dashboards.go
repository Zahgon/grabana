package grabana

import (
	"context"
	"errors"

	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/sdk"
)

// ErrDashboardNotFound is returned when the given dashboard can not be found.
var ErrDashboardNotFound = errors.New("dashboard not found")

// Dashboard represents a Grafana dashboard.
type Dashboard struct {
	ID          int      `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderID    int      `json:"folderId"`
	FolderUID   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderURL   string   `json:"folderUrl"`
}

// GetDashboardByTitle finds a dashboard, given its title.
func (client *Client) GetDashboardByTitle(ctx context.Context, title string) (*Dashboard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rawDashboardByUID finds a dashboard, given its UID.
func (client *Client) rawDashboardByUID(ctx context.Context, uid string) (*sdk.Board, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpsertDashboard creates or replaces a dashboard, in the given folder.
func (client *Client) UpsertDashboard(ctx context.Context, folder *Folder, builder dashboard.Builder) (*Dashboard, error) {
	_ = "STUB: not implemented"
	// first pass: save the new dashboard
	return nil, nil
}

// second pass: delete existing alerts associated to that dashboard

// third pass: create new alerts

// If there are no alerts to create, we can return early

func (client *Client) persistDashboard(ctx context.Context, folder *Folder, builder dashboard.Builder) (*Dashboard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteDashboard deletes a dashboard given its UID.
func (client *Client) DeleteDashboard(ctx context.Context, uid string) error {
	_ = "STUB: not implemented"
	// first: delete existing alerts associated to that dashboard
	return nil
}

// then: delete the dashboard itself

func panelIDByTitle(board *sdk.Board, title string) string { _ = "STUB: not implemented"; return "" }
