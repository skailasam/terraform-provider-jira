package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "jira_user" "foo" {
				  email_address = "roy@klopper.dev"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.jira_user.foo", "id", regexp.MustCompile("^[a-zA-Z0-9]+$")),
					resource.TestMatchResourceAttr("data.jira_user.foo", "email_address", regexp.MustCompile("^roy")),
					resource.TestMatchResourceAttr("data.jira_user.foo", "display_name", regexp.MustCompile("^Roy")),
				),
			},
		},
	})
}
