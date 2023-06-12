package client

import (
	"context"
	"net/http"

	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
)

type NewRelicServices interface {
	ListAlerts(context.Context) (alerts.ListPoliciesParams, *http.Response, error)
}
