package grabana

import (
	"context"
	"errors"

	"github.com/K-Phoen/grabana/datasource"
)

// ErrDatasourceNotFound is returned when the given datasource can not be found.
var ErrDatasourceNotFound = errors.New("datasource not found")

const defaultDatasourceKey = "$grabana_default_datasource_key$"

// UpsertDatasource creates or replaces a datasource.
func (client *Client) UpsertDatasource(ctx context.Context, datasource datasource.Datasource) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteDatasource deletes a datasource given its name.
func (client *Client) DeleteDatasource(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDatasourceUIDByName finds a datasource UID given its name.
func (client *Client) GetDatasourceUIDByName(ctx context.Context, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// datasourcesUIDMap builds a map of datasources UIDs indexed by their name.
func (client *Client) datasourcesUIDMap(ctx context.Context) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getDatasourceIDByName finds a datasource, given its name.
func (client *Client) getDatasourceIDByName(ctx context.Context, name string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
