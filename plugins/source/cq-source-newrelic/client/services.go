package client

import (
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"github.com/newrelic/newrelic-client-go/v2/pkg/plugins"
	"github.com/newrelic/newrelic-client-go/v2/pkg/synthetics"
)

type AlertService interface {
	ListPolicies(*alerts.ListPoliciesParams) ([]alerts.Policy, error)
}

type ApplicationService interface {
	ListApplications(*apm.ListApplicationsParams) ([]*apm.Application, error)
}

type PluginService interface {
	ListPlugins(*plugins.ListPluginsParams) ([]*plugins.Plugin, error)
}

type SyntheticService interface {
	ListMonitors() ([]*synthetics.Monitor, error)
}
