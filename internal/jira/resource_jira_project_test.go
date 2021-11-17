package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProject(t *testing.T) {
	t.Skip("not used")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "jira_user" "foo" {
				  email_address = "test@joob.dev"
				  display_name = "test"
				}
				resource "jira_project" "foo" {
				  name = "test project"
				  key = "TPR"
				  lead = jira_user.foo.id
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("jira_project.foo", "name", regexp.MustCompile("^test project$")),
				),
			},
		},
	})
}
