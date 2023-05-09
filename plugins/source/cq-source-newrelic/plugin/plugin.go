package plugin

import (
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/relic/cq-source-newrelic/client"
	"github.com/relic/cq-source-newrelic/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"relic-newrelic",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
