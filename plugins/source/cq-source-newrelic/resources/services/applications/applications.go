package applications

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"

	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"

	"github.com/newrelic/cq-source-newrelic/client"
	"golang.org/x/sync/errgroup"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:      "new_relic_applications",
		Resolver:  getApplications,
		Transform: transformers.TransformWithStruct(&apm.Application{}),
	}
}

func getApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	apps, err := svc.Services.APM.ListApplications(&apm.ListApplicationsParams{
		Name: "Bewiz",
	})

	fmt.Printf("err: %v+\n", apps)

	if err != nil {
		fmt.Printf("err: %v+\n", apps)
	}

	fmt.Printf("apps: %v+\n", apps)

	res <- apps

	g := errgroup.Group{}
	g.SetLimit(10)

	return g.Wait()
}
