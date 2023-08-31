package services

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cq-source-newrelic/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "newrelic_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: arrow.BinaryTypes.String,
			},
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	return fmt.Errorf("not implemented. client id: " + cl.ID())
}
