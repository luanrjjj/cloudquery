package resources

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/newrelic/cq-source-newrelic/resources/services/alerts"
	"github.com/newrelic/cq-source-newrelic/resources/services/applications"
	"github.com/newrelic/cq-source-newrelic/resources/services/plugins"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		alerts.Alerts(),
		applications.Applications(),
		plugins.Plugins(),
	}
}

// type Comic struct {
// 	ID int `json:"id"`
// }

// func SampleTable() *schema.Table {
// 	return &schema.Table{
// 		Name:      "newrelic_sample_table",
// 		Resolver:  BuildContext,
// 		Transform: transformers.TransformWithStruct(Comic{}),
// 	}
// }

// func getAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

// 	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))

// 	policies, err := client.Alerts.ListPolicies(&alerts.ListPoliciesParams{
// 		Name: "Bewiz",
// 	})
// 	if err != nil {
// 		fmt.Printf("err: %v+\n", policies)
// 	}

// 	fmt.Printf("Policies: %v+\n", policies)

// 	res <- policies

// 	g := errgroup.Group{}
// 	g.SetLimit(10)

// 	return g.Wait()
// }
