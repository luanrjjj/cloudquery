package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/luanrjjj/cq-source-xqcd/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
