package main

import (
	"github.com/buildkite/go-buildkite/buildkite"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: buildkite.Provider,
	})
}
