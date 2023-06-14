package client

import (
	"github.com/newrelic/cq-source-newrelic/client/services"
	"github.com/newrelic/newrelic-client-go/v2/newrelic"
)

type NewRelicServices struct {
	AlertsAPI services.AlertsAPIClient
}

func NewNewRelicServices(apiClient *newrelic.NewRelic) NewRelicServices {
	return NewRelicServices{
		AlertsAPI: newrelic.Alerts.ListPolcies(apiClient),
	}
}
