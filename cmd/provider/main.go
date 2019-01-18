package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return trackerprovider.Create(pt.NewClient)
		},
	})
}
