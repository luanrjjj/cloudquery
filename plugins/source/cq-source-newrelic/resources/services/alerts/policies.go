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

func Policies() *schema.Table {
	return &schema.Table{
		Name:     "new_relic_alert_policies",
		Resolver: getPolicies,
		//Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&alerts.Policy{}),
	}
}

func getPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	policies, err := svc.Services.Alert.ListPolicies(&alerts.ListPoliciesParams{})

	if err != nil {
		fmt.Printf("errHAHAHAAH: %v+\n", err)
	}

	res <- policies

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
