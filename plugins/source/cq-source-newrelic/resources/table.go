package resources

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

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:      "newrelic_sample_table",
		Resolver:  BuildContext,
		Transform: transformers.TransformWithStruct(Comic{}),
	}
}

func getAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

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

func BuildContext(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	// c := meta
	// fmt.Println("ushduasdhauhduasdhsauh", os.Getenv("NEW_RELIC_API_KEY"))
	// value, err := newrelic.New(newrelic.ConfigPersonalAPIKey(os.Getenv("NEW_RELIC_API_KEY")))
	// // fmt.Println("ushduasdhauhduasdhsauh", err)
	// if err != nil {
	// 	// ...
	// 	fmt.Println("error", err)
	// }
	// queryBuilder := entities.EntitySearchQueryBuilder{
	// 	Name: "Bewiz",
	// 	Type: entities.EntitySearchQueryBuilderTypeTypes.DASHBOARD,
	// }

	// entitySearch, err := value.Entities.GetEntitySearch(
	// 	entities.EntitySearchOptions{},
	// 	"",
	// 	queryBuilder,
	// 	[]entities.EntitySearchSortCriteria{},
	// )

	// fmt.Println("entitySearch", entitySearch)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// apps, err := value.APM.ListApplications(&apm.ListApplicationsParams{
	// 	Name: "Bewiz",
	// })

	// fmt.Printf("apps", apps)

	// res <- entitySearch.Results.Entities

	// g := errgroup.Group{}
	// g.SetLimit(10)

	// return g.Wait()

	return getAlerts(ctx, meta, parent, res)
}
