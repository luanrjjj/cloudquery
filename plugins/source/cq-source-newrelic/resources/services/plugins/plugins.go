package plugins

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/pkg/plugins"

	"github.com/newrelic/cq-source-newrelic/client"
	"golang.org/x/sync/errgroup"
)

func Plugins() *schema.Table {
	return &schema.Table{
		Name:      "new_relic_plugins",
		Resolver:  getPlugins,
		Transform: transformers.TransformWithStruct(&plugins.Plugin{}),
	}
}

func getPlugins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	plugins, err := svc.Services.Plugin.ListPlugins(&plugins.ListPluginsParams{})

	fmt.Printf("err: %v+\n", plugins)

	if err != nil {
		fmt.Printf("err: %v+\n", plugins)
	}

	fmt.Printf("plugins: %v+\n", plugins)

	res <- plugins

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
