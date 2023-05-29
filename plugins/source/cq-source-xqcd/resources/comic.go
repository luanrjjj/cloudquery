// package resources

// import (
// 	"context"
// 	"fmt"

// 	"github.com/cloudquery/plugin-sdk/schema"
// )

// func SampleTable() *schema.Table {
// 	return &schema.Table{
// 		Name:     "xqcd_sample_table",
// 		Resolver: fetchSampleTable,
// 		Columns: []schema.Column{
// 			{
// 				Name: "column",
// 				Type: schema.TypeString,
// 			},
// 		},
// 	}
// }

// func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
// 	return fmt.Errorf("not implemented")
// }

package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/luanrjjj/cq-source-xqcd/client"
	"github.com/luanrjjj/cq-source-xqcd/internal/xkcd"
	"golang.org/x/sync/errgroup"
)

func ComicsTable() *schema.Table {
	return &schema.Table{
		Name:      "xkcd_comics",
		Resolver:  fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
	}
}

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	c.Logger.Info().Msgf("AUSDHAUSHDUASDHUSAHDUASHDUASHDUASHDASUHDSUA: %v", c)
	c.Logger.Info().Msgf("HAHAHAHAAHAHAH: %v", c.XKCD)
	comic, err := c.XKCD.GetLatestComic(0)
	c.Logger.Info().Msgf("comic: %v", err)
	if err != nil {
		return err
	}
	res <- comic
	g := errgroup.Group{}
	g.SetLimit(10)
	for i := 1; i < comic.Num; i++ {
		i := i
		g.Go(func() error {
			comic, err := c.XKCD.GetComic(i)
			if err != nil {
				c.Logger.Error().Err(err).Msgf("failed to fetch comic %d", i)
				return err
			}
			res <- comic
			return nil
		})
	}
	return g.Wait()
}
