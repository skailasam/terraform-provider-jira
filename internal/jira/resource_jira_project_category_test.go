package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProjectCategory(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "jira_project_category" "foo" {
				  name = "test project"
				  description = "Some description"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("jira_project_category.foo", "id", regexp.MustCompile("^[a-zA-Z0-9]+$")),
					resource.TestMatchResourceAttr("jira_project_category.foo", "name", regexp.MustCompile("^test project$")),
					resource.TestMatchResourceAttr("jira_project_category.foo", "description", regexp.MustCompile("^Some description$")),
				),
			},
		},
	})
}
