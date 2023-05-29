package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/luanrjjj/cq-source-xqcd/client"
	"github.com/luanrjjj/cq-source-xqcd/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"luanrjjj-xqcd",
		Version,
		schema.Tables{
			resources.ComicsTable(),
		},
		client.New,
	)
}
