package buildkite

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccResourcePipelineCreate = `
resource buildkite_pipeline "test" {
  name         = "Acceptance test :terraform: Pipeline"
  organization = "cozero"
  description  = "Generated via acceptance tests - please delete if left dangling"
  repository   = "git@github.com:COzero/terraform-provider-buildkite.git"

  step {
    type    = "script"
    name    = "Hi!"
    command = "echo \"Hello world\""
  }
}
`

func TestAccResourcePipelineCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourcePipelineCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("buildkite_pipeline.test", "web_url"),
				),
			},
		},
	})
	return
}
