package buildkite

import (
	"fmt"
	"log"
	"time"

	"github.com/buildkite/go-buildkite/buildkite"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		// Ref https://buildkite.com/docs/rest-api/pipelines#create-a-pipeline
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Schema: map[string]*schema.Schema{
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateStepType(),
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"command": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"env": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"timeout_in_minutes": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"agent_query_rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"artifact_paths": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"branch_configuration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"concurrency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parallelism": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"provider_settings": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"branch_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"skip_queued_branch_builds": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"skip_queued_branch_builds_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cancel_running_branch_builds": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"cancel_running_branch_builds_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"team_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_branch": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhook_url": &schema.Schema{
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

func resourcePipelineCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*buildkite.Client)

	input := buildPipelineInput(d)

	pipe, _, err := client.Pipelines.Create(d.Get("organization").(string), input)

	if err != nil {
		return fmt.Errorf("Error creating pipeline: %s", err)
	}

	updatePipelineFromAPI(d, pipe)

	return resourcePipelineRead(d, meta)
}

func resourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*buildkite.Client)

	pipe, resp, err := client.Pipelines.Get(d.Get("organization").(string), d.Get("slug").(string))
	if err != nil {
		// If pipeline not found, delete if from the state file
		if resp.StatusCode == 404 {
			log.Printf("[WARN] Pipeline (%s) not found, removing from state", d.Get("slug"))
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading pipeline: %s", err)
	}

	updatePipelineFromAPI(d, pipe)

	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*buildkite.Client)

	pipe := &buildkite.Pipeline{
		Name:       String(d.Get("name").(string)),
		Repository: String(d.Get("repository").(string)),
	}

	steps := d.Get("step").(*schema.Set).List()
	pipe.Steps = make([]*buildkite.Step, len(steps))

	for i, s := range steps {
		step := s.(map[string]interface{})
		pipe.Steps[i] = &buildkite.Step{
			Type:                String(step["type"].(string)),
			Name:                String(step["name"].(string)),
			Command:             String(step["command"].(string)),
			ArtifactPaths:       String(step["artifact_paths"].(string)),
			BranchConfiguration: String(step["branch_configuration"].(string)),
			Env:                 map[string]string{},
			TimeoutInMinutes:    step["timeout_in_minutes"].(int),
			// Not yet supported by API client
			// Concurrency:         step["concurrency"].(int),
			// Parallelism:         step["parallelism"].(int),
		}

		for k, v := range step["env"].(map[string]interface{}) {
			pipe.Steps[i].Env[k] = v.(string)
		}

		agentQueryRules := make([]string, len(step["agent_query_rules"].([]interface{})))

		for j, v := range step["agent_query_rules"].([]interface{}) {
			agentQueryRules[j] = v.(string)
		}

		pipe.Steps[i].AgentQueryRules = agentQueryRules
	}

	client.Pipelines.Update(d.Get("organization").(string), pipe)

	updatePipelineFromAPI(d, pipe)

	return resourcePipelineRead(d, meta)
}

func resourcePipelineDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*buildkite.Client)

	_, err := client.Pipelines.Delete(d.Get("organization").(string), d.Get("slug").(string))
	if err != nil {
		return fmt.Errorf("Error deleting pipeline: %s", err)
	}

	return nil
}

func buildPipelineInput(d *schema.ResourceData) *buildkite.CreatePipeline {
	input := &buildkite.CreatePipeline{
		Name:                            d.Get("name").(string),
		Repository:                      d.Get("repository").(string),
		Description:                     d.Get("description").(string),
		Env:                             map[string]string{},
		ProviderSettings:                map[string]bool{},
		BranchConfiguration:             d.Get("branch_configuration").(string),
		SkipQueuedBranchBuilds:          d.Get("skip_queued_branch_builds").(bool),
		SkipQueuedBranchBuildsFilter:    d.Get("skip_queued_branch_builds_filter").(string),
		CancelRunningBranchBuilds:       d.Get("cancel_running_branch_builds").(bool),
		CancelRunningBranchBuildsFilter: d.Get("cancel_running_branch_builds_filter").(string),
		TeamUuids:                       []string{},
	}

	steps := d.Get("step").(*schema.Set).List()
	input.Steps = make([]buildkite.Step, len(steps))

	for i, s := range steps {
		step := s.(map[string]interface{})
		input.Steps[i] = buildkite.Step{
			Type:                String(step["type"].(string)),
			Name:                String(step["name"].(string)),
			Command:             String(step["command"].(string)),
			ArtifactPaths:       String(step["artifact_paths"].(string)),
			BranchConfiguration: String(step["branch_configuration"].(string)),
			Env:                 map[string]string{},
			TimeoutInMinutes:    step["timeout_in_minutes"].(int),
			// Not yet supported by API client
			// Concurrency:         step["concurrency"].(int),
			// Parallelism:         step["parallelism"].(int),
		}

		for k, v := range step["env"].(map[string]interface{}) {
			input.Steps[i].Env[k] = v.(string)
		}

		agentQueryRules := make([]string, len(step["agent_query_rules"].([]interface{})))

		for j, v := range step["agent_query_rules"].([]interface{}) {
			agentQueryRules[j] = v.(string)
		}

		input.Steps[i].AgentQueryRules = agentQueryRules
	}

	for k, v := range d.Get("env").(map[string]interface{}) {
		input.Env[k] = v.(string)
	}

	for k, v := range d.Get("provider_settings").(map[string]interface{}) {
		input.ProviderSettings[k] = v.(bool)
	}

	for i, v := range d.Get("team_uuids").([]interface{}) {
		input.TeamUuids[i] = v.(string)
	}

	return input
}

func updatePipelineFromAPI(d *schema.ResourceData, p *buildkite.Pipeline) {
	d.SetId(StringValue(p.ID))
	d.Set("name", StringValue(p.Name))
	d.Set("badge_url", StringValue(p.BadgeURL))
	d.Set("builds_url", StringValue(p.BuildsURL))
	d.Set("created_at", p.CreatedAt.Format(time.RFC3339))
	d.Set("webhook_url", StringValue(p.Provider.WebhookURL))
	d.Set("repository", StringValue(p.Repository))
	d.Set("slug", StringValue(p.Slug))
	d.Set("url", StringValue(p.URL))
	d.Set("web_url", StringValue(p.WebURL))

	steps := buildStepsFromAPI(p.Steps)
	d.Set("steps", steps)

	return
}

func buildStepsFromAPI(steps []*buildkite.Step) []interface{} {
	out := make([]interface{}, len(steps))
	for i, step := range steps {
		out[i] = map[string]interface{}{
			"agent_query_rules":    step.AgentQueryRules,
			"artifact_paths":       step.ArtifactPaths,
			"branch_configuration": step.BranchConfiguration,
			"command":              step.Command,
			"env":                  step.Env,
			"name":                 step.Name,
			"timeout_in_minutes":   step.TimeoutInMinutes,
			"type":                 step.Type,
		}
	}

	return out
}
