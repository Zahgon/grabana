package grabana

import (
	"context"
	"errors"
)

// ErrFolderNotFound is returned when the given folder can not be found.
var ErrFolderNotFound = errors.New("folder not found")

// Folder represents a dashboard folder.
// See https://grafana.com/docs/grafana/latest/reference/dashboard_folders/
type Folder struct {
	ID        uint   `json:"id"`
	UID       string `json:"uid"`
	ParentUID string `json:"parentUid"`
	Title     string `json:"title"`
}

// FindOrCreateFolder returns the folder by its name or creates it if it doesn't exist.
func (client *Client) FindOrCreateFolder(ctx context.Context, name string) (*Folder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateFolder creates a dashboard folder.
// See https://grafana.com/docs/grafana/latest/reference/dashboard_folders/
func (client *Client) CreateFolder(ctx context.Context, name string) (*Folder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFolderByTitle finds a folder, given its title.
func (client *Client) GetFolderByTitle(ctx context.Context, title string) (*Folder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
