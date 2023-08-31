package alerts

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"

	"golang.org/x/sync/errgroup"
)

type Comic struct {
	ID int `json:"id"`
}

func Alerts() *schema.Table {
	return &schema.Table{
		Name:     "datadog_alerts",
		Resolver: getAlerts,
		// Multiplex: client.AccountMultiplex,
		// Resolver:  BuildContext,
		Transform: transformers.TransformWithStruct(Comic{}),
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

	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))

	if err != nil {
		fmt.Printf("err: %v+\n", client)
	}

	policies, err := client.Alerts.ListPolicies(&alerts.ListPoliciesParams{
		Name: "Example policy",
	})
	if err != nil {
		fmt.Printf("err: %v+\n", policies)
	}

	fmt.Printf("Policies: %v+\n", policies)

	res <- policies

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
