package buildkite

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func validateStepType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"manual",
		"script",
		"trigger",
		"waiter",
	}, false)
}
