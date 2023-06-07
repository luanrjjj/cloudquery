package main

import (
	"github.com/newrelic/cq-source-newrelic/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
