package buildkite

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// We need a pipeline resource so that this can be tested properly
const testAccDataSourcePipelineRead = `
data buildkite_pipeline "test" {
  name = "Stu's Great Showcase Pipeline"
  organization = "cozero"
}
`

func TestAccDataSourcePipelineRead(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourcePipelineRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.buildkite_pipeline.test", "web_url"),
				),
			},
		},
	})
	return
}
