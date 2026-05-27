package grabana

import (
	"context"
	"errors"

	"github.com/K-Phoen/grabana/alert"
	"github.com/K-Phoen/grabana/alertmanager"
)

// ErrAlertNotFound is returned when the requested alert can not be found.
var ErrAlertNotFound = errors.New("alert not found")

type alertRef struct {
	Namespace string
	RuleGroup string
}

// ConfigureAlertManager updates the alert manager configuration.
func (client *Client) ConfigureAlertManager(ctx context.Context, manager *alertmanager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// AddAlert creates an alert group within a given namespace.
func (client *Client) AddAlert(ctx context.Context, namespace string, alertDefinition alert.Alert, datasourcesMap map[string]string) error {
	_ = "STUB: not implemented"
	// Find out which datasource the alert depends on, and inject its UID into the sdk definition
	return nil
}

// Before we can add this alert, we need to delete any other alert that might exist for this dashboard and panel

// Save the alert!

// DeleteAlertGroup deletes an alert group.
func (client *Client) DeleteAlertGroup(ctx context.Context, namespace string, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// listAlertsForDashboard fetches a list of alerts linked to the given dashboard.
func (client *Client) listAlertsForDashboard(ctx context.Context, dashboardUID string) ([]alertRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
