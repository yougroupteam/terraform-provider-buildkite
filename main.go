package main

import (
	"github.com/cozero/terraform-provider-buildkite/buildkite"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: buildkite.Provider,
	})
}
