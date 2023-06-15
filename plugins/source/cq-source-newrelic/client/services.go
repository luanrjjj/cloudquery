package client

import (
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
)

type AlertService interface {
	ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
	CreatePolicy(alerts.Policy) (*alerts.Policy, error)
}
