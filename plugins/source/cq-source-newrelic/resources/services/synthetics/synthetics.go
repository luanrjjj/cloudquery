package synthetics

import (
	"context"
	"fmt"
	"github.com/apache/arrow/go/v13/arrow"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/pkg/plugins"

	"github.com/newrelic/cq-source-newrelic/client"
	"golang.org/x/sync/errgroup"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "new_relic_synthetics",
		Resolver:  getSynthetics,
		Transform: transformers.TransformWithStruct(&plugins.Plugin{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}

func getSynthetics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	apps, err := svc.Services.Synthetic.ListMonitors()

	if err != nil {
		fmt.Printf("err: %v+\n", apps)
	}

	fmt.Printf("apps: %v+\n", apps)

	res <- apps

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
