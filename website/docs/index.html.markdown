---
layout: "buildkite"
page_title: "Provider: Buildkite"
sidebar_current: "docs-buildkite-index"
description: |-
  The Buildkite provider is used to interact with the resources supported by Buildkite. The provider needs to be configured with the proper credentials before it can be used.
---

# Buildkite Provider

The Buildkite provider is used to interact with the resources supported by
Buildkite. The provider needs to be configured with the proper credentials
before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Buildkite Provider
provider "buildkite" {
  api_token = "${var.buildkite_api_token}"
}

# Create a pipeline
resource "buildkite_pipeline" "legendary" {
  # ...
}
```

## Authentication

You can obtain an API Access Token from https://buildkite.com/user/api-access-tokens

## Argument Reference

The following arguments are supported in the `provider` block:

* `api_token` - (Required) This is the Buildkite API Access Token.
