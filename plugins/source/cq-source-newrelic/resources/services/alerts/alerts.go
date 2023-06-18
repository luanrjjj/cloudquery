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

func Alerts() *schema.Table {
	return &schema.Table{
		Name:     "new_relic_alerts",
		Resolver: getAlerts,
		// Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&alerts.Policy{}),
		// Transform: transformers.TransformWithStruct(&datadogV2.IncidentResponseData{}),
		// Columns: []schema.Column{
		// 	{
		// 		Name:       "account_name",
		// 		Type:       arrow.BinaryTypes.String,
		// 		Resolver:   client.ResolveAccountName,
		// 		PrimaryKey: true,
		// 	},
		// 	{
		// 		Name:       "id",
		// 		Type:       arrow.BinaryTypes.String,
		// 		Resolver:   schema.PathResolver("Id"),
		// 		PrimaryKey: true,
		// 	},
		// },
	}
}

func getAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	policies, err := svc.Services.Alert.ListPolicies(&alerts.ListPoliciesParams{
		Name: "Bewiz",
	})

	fmt.Printf("err: %v+\n", policies)

	if err != nil {
		fmt.Printf("err: %v+\n", policies)
	}

	fmt.Printf("Policies: %v+\n", policies)

	res <- policies

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
