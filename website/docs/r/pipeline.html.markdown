---
layout: "buildkite"
page_title: "Buildkite: buildkite_pipeline"
sidebar_current: "docs-buildkite-resource-pipeline"
description: |-
  Provides a Buildkite Pipeline
---

# Resource: buildkite_pipeline

Provides a Buildkite Pipeline

## Example Usage

```hcl
resource "buildkite_pipeline" "example" {
  organization = "my-org"
  name         = "My great pipeline"
  repository   = "git@github.com:myorg/my-repo.git"
  step {
    type    = "script"
    name    = "Hi!"
    command = "echo \"Hello world\""
  }
}
```

## Argument Reference

 * `organization` - (Required) The name of the organization to which the pipeline is associated
 * `name` - (Required) The pipeline name
 * `repository` - (Required) The git repository associated with the pipeline
 * `step` - (Required) At least one step block as documented below
 * `description` - (Optional, Forces new resource) The pipeline description
 * `env` - (Optional, Forces new resource) A map of ENV key/value pairs that apply to the whole pipeline
 * `provider_settings` - (Optional, Forces new resource) A map of key/bool provider_settings. Provider specific and are not well documented.
 * `branch_configuration` - (Optional, Forces new resource) Sets the default branch?
 * `skip_queued_branch_builds` - (Optional, Forces new resource) When a new build is created on a branch, any previous builds that haven't yet started on the same branch will be automatically marked as skipped.
 * `skip_queued_branch_builds_filter` - (Optional, Forces new resource) Limit which branches build skipping applies to, for example !master will ensure that the master branch won't have it's builds automatically skipped.
 * `cancel_running_branch_builds` - (Optional, Forces new resource) When a new build is created on a branch, any previous builds that are running on the same branch will be automatically cancelled.
 * `cancel_running_branch_builds_filter` - (Optional, Forces new resource) Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
 * `team_uuids` - (Optional, Forces new resource) A list of team UUIDs that have access to the pipeline.

## Step
 * `type` - (Required) The type of step. Valid values are manual, script, trigger or waiter
 * `name` - (Optional) The name of the step
 * `command` - (Optional) - The command to run
 * `env` - (Optional) A map of ENV key/value pairs that apply to the step
 * `timeout_in_minutes` - (Optional) Number of minutes the step can run for before being automatically canceled.
 * `agent_query_rules` - (Optional) A list of query rules for selecting an agent, e.g. `ruby=true`
 * `artifact_paths` - (Optional) File paths for the artifacts to be automatically uploaded after the step has run (e.g. tmp/\*\*/\*.png). Separate each path with a semi-colon.
 * `branch_configuration` - (Optional) List of branch conditions, each separated with a space. e.g. master. This step will be skipped if the build’s branch does not match.

## Attributes Reference

 * `badge_url` - The URL of a badge that displays current pipeline build status
 * `builds_url` - The API endpoint to access builds for this pipeline
 * `created_at` - creation time for the pipeline (RFC3339)
 * `default_branch` - The default development branch for this repository.
 * `provider_id` - The git provider (e.g. github, bitbucket) associated with the pipeline
 * `slug` - The slug of the pipeline
 * `url` - The API endpoint for this pipeline
 * `web_url` - The web endpoint for this pipeline
 * `webhook_url` - The URL at which you should fire build webhooks
