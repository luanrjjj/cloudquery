package services

import (
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
)

type AlertsAPIClient interface {
	ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
	CreatePolicy(alerts.Policy) (*alerts.Policy, error)
}
