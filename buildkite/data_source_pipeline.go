package buildkite

import (
	"fmt"
	"time"

	"github.com/buildkite/go-buildkite/buildkite"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePipeline() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePipelineRead,

		Schema: map[string]*schema.Schema{
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"builds_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"badge_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*buildkite.Client)

	p, _, err := client.Pipelines.Get(d.Get("organization").(string), d.Get("slug").(string))
	if err != nil {
		return fmt.Errorf("Error reading pipeline: %s", err)
	}

	d.SetId(StringValue(p.ID))

	d.Set("badge_url", StringValue(p.BadgeURL))
	d.Set("builds_url", StringValue(p.BuildsURL))
	d.Set("created_at", p.CreatedAt.Format(time.RFC3339))
	d.Set("name", StringValue(p.Name))
	d.Set("provider_id", StringValue(p.Provider.ID))
	d.Set("repository", StringValue(p.Repository))
	d.Set("slug", StringValue(p.Slug))
	d.Set("url", StringValue(p.URL))
	d.Set("web_url", StringValue(p.WebURL))
	d.Set("webhook_url", StringValue(p.Provider.WebhookURL))

	steps := buildStepsFromAPI(p.Steps)
	d.Set("steps", steps)

	return nil
}
