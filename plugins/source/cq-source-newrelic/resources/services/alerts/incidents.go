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

func Incidents() *schema.Table {
	return &schema.Table{
		Name:      "new_relic_alert_incidents",
		Resolver:  getIncidents,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&alerts.Incident{}),
		//	Columns: []schema.Column{
		//		{
		//			Name:       "account_name",
		//			Type:       arrow.BinaryTypes.String,
		//			Resolver:   client.ResolveAccountName,
		//			PrimaryKey: true,
		//		},
		//		{
		//			Name:       "id",
		//			Type:       arrow.BinaryTypes.String,
		//			Resolver:   schema.PathResolver("Id"),
		//			PrimaryKey: true,
		//		},
		//	},
	}
}

func getIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	incidents, err := svc.Services.Alert.ListIncidents(false, false)

	if err != nil {
		fmt.Printf("err: %v+\n", incidents)
	}

	fmt.Printf("Policies: %v+\n", incidents)

	res <- incidents

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
