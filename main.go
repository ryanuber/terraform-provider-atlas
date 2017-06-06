package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-atlas/atlas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: atlas.Provider})
}
