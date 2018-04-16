package buildkite

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"buildkite_pipeline": dataSourcePipeline(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"buildkite_pipeline": resourcePipeline(),
		},

		Schema: map[string]*schema.Schema{
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUILDKITE_API_TOKEN", nil),
			},
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("api_token").(string),
	}

	meta, err := config.Client()
	if err != nil {
		return nil, err
	}

	return meta, nil
}
