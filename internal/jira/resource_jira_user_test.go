package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
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
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("jira_user.foo", "id", regexp.MustCompile("^[a-zA-Z0-9]+$")),
					resource.TestMatchResourceAttr("jira_user.foo", "email_address", regexp.MustCompile("^test")),
					resource.TestMatchResourceAttr("jira_user.foo", "display_name", regexp.MustCompile("^test$")),
				),
			},
		},
	})
}
