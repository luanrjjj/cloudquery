package main

import (
	"github.com/cloudquery/plugin-sdk/v2/serve"
	"github.com/relic/cq-source-newrelic/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
