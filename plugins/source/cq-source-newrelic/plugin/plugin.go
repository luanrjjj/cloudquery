package plugin

import (
	"github.com/newrelic/cq-source-newrelic/client"
	"github.com/newrelic/cq-source-newrelic/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"newrelic-newrelic",
		Version,
		resources.Tables(),
		client.New,
	)
}
