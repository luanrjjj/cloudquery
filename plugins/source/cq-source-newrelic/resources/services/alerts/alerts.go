package alerts

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"

	"golang.org/x/sync/errgroup"
)

type Comic struct {
	ID int `json:"id"`
}

func alerts() *schema.Table {
	return &schema.Table{
		Name:     "datadog_alerts",
		Resolver: fetchalerts,
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

func fetchalerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.alertsAPI.Listalerts(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}

func getAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	os.Setenv("NEW_RELIC_API_KEY", "NRAK-SMYODD87CECC7JNKN8JVKZRM32P")

	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))

	policies, err := client.Alerts.ListPolicies(&alerts.ListPoliciesParams{
		Name: "Bewiz",
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
