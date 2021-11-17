package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIssueType(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "jira_issue_type" "foo" {
				  name = "test project"
				  description = "Some description"
				  hierarchy_level = 0
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("jira_issue_type.foo", "id", regexp.MustCompile("^[a-zA-Z0-9]+$")),
					resource.TestMatchResourceAttr("jira_issue_type.foo", "name", regexp.MustCompile("^test project$")),
					resource.TestMatchResourceAttr("jira_issue_type.foo", "description", regexp.MustCompile("^Some description$")),
				),
			},
		},
	})
}
