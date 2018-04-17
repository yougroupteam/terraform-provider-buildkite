---
layout: "buildkite"
page_title: "Buildkite: buildkite_pipeline"
sidebar_current: "docs-buildkite-datasource-pipeline"
description: |-
  Get information on a Buildkite Pipeline
---

# Data Source: buildkite_pipeline

Use this data source to get information about a Buildkite pipeline.

## Example Usage

```hcl
data "buildkite_pipeline" "example" {
  organization = "my-org"
  slug         = "mega-pipeline"
}
```

## Argument Reference

 * `organization` - (Required) The name of the organization to which the pipeline is associated
 * `slug` - (Required) The slug of the pipeline to look up. If no pipeline is found with this slug, an error will be returned.

## Attributes Reference

 * `badge_url` - The URL of a badge that displays current pipeline build status
 * `builds_url` - The API endpoint to access builds for this pipeline
 * `created_at` - creation time for the pipeline (RFC3339)
 * `name` - The pipeline name
 * `provider_id` - The git provider (e.g. github, bitbucket) associated with the pipeline
 * `repository` - The git repository associated with the pipeline
 * `url` - The API endpoint for this pipeline
 * `web_url` - The web endpoint for this pipeline
 * `webhook_url` - The URL at which you should fire build webhooks
