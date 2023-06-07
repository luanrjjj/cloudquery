package plugin

import (
	"github.com/newrelic/cq-source-newrelic/client"
	"github.com/newrelic/cq-source-newrelic/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"newrelic-newrelic",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
