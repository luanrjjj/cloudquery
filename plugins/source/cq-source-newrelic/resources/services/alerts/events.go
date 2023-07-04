package alerts

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"

	"github.com/newrelic/cq-source-newrelic/client"
	"golang.org/x/sync/errgroup"
)

func Events() *schema.Table {
	return &schema.Table{
		Name:     "new_relic_events",
		Resolver: getEvents,
		//Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&alerts.AlertEvent{}),
		//Columns: []schema.Column{
		//	{
		//		Name:       "account_name",
		//		Type:       arrow.BinaryTypes.String,
		//		Resolver:   client.ResolveAccountName,
		//		PrimaryKey: true,
		//	},
		//	{
		//		Name:       "id",
		//		Type:       arrow.BinaryTypes.String,
		//		Resolver:   schema.PathResolver("Id"),
		//		PrimaryKey: true,
		//	},
		//},
	}
}

func getEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	events, err := svc.Services.Alert.ListAlertEvents(&alerts.ListAlertEventsParams{})

	fmt.Printf("err: %v+\n", events)

	if err != nil {
		fmt.Printf("err: %v+\n", events)
	}

	fmt.Printf("Policies: %v+\n", events)

	res <- events

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
