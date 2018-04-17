package buildkite

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourcePipelineRead(t *testing.T) {
	rStr := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourcePipeline(rStr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.buildkite_pipeline.test",
						"name",
						fmt.Sprintf("Acceptance test :terraform: %s", rStr),
					),
				),
			},
		},
	})
	return
}

func testAccDataSourcePipeline(rStr string) string {
	return fmt.Sprintf(`
resource buildkite_pipeline "test" {
  name         = "Acceptance test :terraform: %s"
  organization = "cozero"
  description  = "Generated via acceptance tests - please delete if left dangling"
  repository   = "git@github.com:COzero/terraform-provider-buildkite.git"

  step {
    type    = "script"
    name    = "Hi!"
    command = "echo \"Hello world\""
  }
}

data buildkite_pipeline "test" {
	organization = "cozero"
	slug         = "${buildkite_pipeline.test.slug}"
}
`, rStr)
}
