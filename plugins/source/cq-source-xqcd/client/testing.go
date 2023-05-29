package client

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/luanrjjj/cq-source-xqcd/internal/xkcd"
	"github.com/rs/zerolog"
)

func TestHelper(t *testing.T, table *schema.Table, ts *httptest.Server) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, opts source.Options) (schema.ClientMeta, error) {
		c, err := xkcd.NewClient(xkcd.WithHTTPClient(ts.Client()), xkcd.WithBaseURL(ts.URL))
		if c != nil {
			fmt.Printf("c: %v", c)
		}
		if err != nil {
			return nil, err
		}
		return &Client{
			Logger: l,
		}, nil
	}
	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
