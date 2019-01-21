package trackerprovider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider/resources/projects"
)

type ProviderClient func(string) pt.ClientCaller

func Create(providerClient ProviderClient) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PVTL_TRACKER_TOKEN", ""),
				Description: "Pivotal Tracker API access token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pivotaltracker_project": projects.NewProjectResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return providerClient(d.Get("access_token").(string)), nil
		},
	}
}
