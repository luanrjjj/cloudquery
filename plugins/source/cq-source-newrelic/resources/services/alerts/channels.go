package alerts

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"

	"github.com/newrelic/cq-source-newrelic/client"
	"golang.org/x/sync/errgroup"
)

func Channels() *schema.Table {
	return &schema.Table{
		Name:      "new_relic_channels",
		Resolver:  getChannels,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&alerts.Channel{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}

func getChannels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	channels, err := svc.Services.Alert.ListChannels()

	fmt.Printf("err: %v+\n", channels)

	if err != nil {
		fmt.Printf("err: %v+\n", channels)
	}

	fmt.Printf("Policies: %v+\n", channels)

	res <- channels

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
