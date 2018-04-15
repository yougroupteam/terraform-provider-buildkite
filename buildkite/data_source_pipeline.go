package buildkite

import (
	"fmt"

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
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*buildkite.Client)
	name := d.Get("name").(string)
	var pipeline buildkite.Pipeline

	// The API is paginated; loop through until we find what we're looking for
	for i, done := 0, false; !done; i++ {
		input := &buildkite.PipelineListOptions{
			ListOptions: buildkite.ListOptions{
				Page: i,
			},
		}
		pipelines, _, err := conn.Pipelines.List(d.Get("organization").(string), input)
		if err != nil {
			return err
		}

		if len(pipelines) == 0 {
			return fmt.Errorf("No pipeline %s found", name)
		}

		for _, p := range pipelines {
			if *p.Name == name {
				pipeline = p
				done = true
				break
			}
		}
	}

	d.SetId(*pipeline.ID)
	d.Set("web_url", pipeline.WebURL)

	return nil
}
