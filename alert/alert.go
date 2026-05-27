package alert

import (
	"github.com/K-Phoen/sdk"
)

// ErrorMode represents the behavior of an alert in case of execution error.
type ErrorMode string

// Alerting will set the alert state to "alerting".
const ErrorAlerting ErrorMode = "Alerting"

// LastState will set the alert state to its previous state.
const ErrorKO ErrorMode = "Error"

// LastState will set the alert state to its previous state.
const ErrorOK ErrorMode = "OK"

// NoDataMode represents the behavior of an alert when no data is returned by
// the query.
type NoDataMode string

// NoData will set the alert state to "no data".
const NoDataEmpty NoDataMode = "NoData"

// Error will set the alert state to "alerting".
const NoDataAlerting NoDataMode = "Alerting"

// OK will set the alert state to "ok".
const NoDataOK NoDataMode = "OK"

// Option represents an option that can be used to configure an alert.
type Option func(alert *Alert)

// Channel represents an alert notification channel.
// See https://grafana.com/docs/grafana/latest/alerting/notifications/#notification-channel-setup
type Channel struct {
	ID   uint   `json:"id"`
	UID  string `json:"uid"`
	Name string `json:"Name"`
	Type string `json:"type"`
}

const alertConditionRef = "_alert_condition_"

// Alert represents an alert that can be triggered by a query.
type Alert struct {
	Builder *sdk.Alert

	// For internal use only
	Datasource   string
	DashboardUID string
	PanelID      string
}

// New creates a new alert.
func New(name string, options ...Option) *Alert { _ = "STUB: not implemented"; return nil }

func defaults() []Option { _ = "STUB: not implemented"; return nil }

func (alert *Alert) HookDatasourceUID(uid string) { _ = "STUB: not implemented"; return }

func (alert *Alert) HookDashboardUID(uid string) { _ = "STUB: not implemented"; return }

func (alert *Alert) HookPanelID(id string) { _ = "STUB: not implemented"; return }

// Summary sets the summary associated to the alert.
func Summary(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Description sets the description associated to the alert.
func Description(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Runbook sets the runbook URL associated to the alert.
func Runbook(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// For sets the time interval during which a query violating the threshold
// before the alert being actually triggered.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#for
func For(duration string) Option { _ = "STUB: not implemented"; return *new(Option) }

// EvaluateEvery defines the evaluation interval.
func EvaluateEvery(interval string) Option { _ = "STUB: not implemented"; return *new(Option) }

// OnExecutionError defines the behavior on execution error.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#execution-errors-or-timeouts
func OnExecutionError(mode ErrorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// OnNoData defines the behavior when the query returns no data.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#no-data-null-values
func OnNoData(mode NoDataMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// If defines a single condition that will trigger the alert.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#conditions
func If(reducer QueryReducer, queryRef string, evaluator ConditionEvaluator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// IfOr defines a single condition that will trigger the alert.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#conditions
func IfOr(reducer QueryReducer, queryRef string, evaluator ConditionEvaluator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ifOperand(operand Operator, reducer QueryReducer, queryRef string, evaluator ConditionEvaluator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Tags defines a set of tags that will be forwarded to the notifications
// channels when the alert will tbe triggered or used to route the alert.
func Tags(tags map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }
