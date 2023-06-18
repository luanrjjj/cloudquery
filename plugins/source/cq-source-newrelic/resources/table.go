package resources

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/newrelic/cq-source-newrelic/resources/services/alerts"
	"github.com/newrelic/cq-source-newrelic/resources/services/applications"
	"github.com/newrelic/cq-source-newrelic/resources/services/plugins"
	"github.com/newrelic/cq-source-newrelic/resources/services/synthetics"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		alerts.Alerts(),
		applications.Applications(),
		plugins.Plugins(),
		synthetics.Synthetics(),
	}
}
