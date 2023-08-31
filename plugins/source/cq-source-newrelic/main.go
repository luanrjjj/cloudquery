package main

import (
<<<<<<< Updated upstream
	"github.com/newrelic/cq-source-newrelic/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
=======
	"context"
	"log"

	"github.com/cloudquery/newrelic/resources/plugin"

	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	if err := serve.Plugin(plugin.Plugin()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
>>>>>>> Stashed changes
}
