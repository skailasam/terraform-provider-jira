package jira

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGroupMembership(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "jira_group" "foo" {
				  name = "test-group"
				}
				resource "jira_user" "foo" {
				  email_address = "foo@joob.dev"
				  display_name = "foo"
				}
				resource "jira_group_membership" "foo" {
				  group_name = jira_group.foo.name
				  account_id = jira_user.foo.id
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("jira_group_membership.foo", "group_name", regexp.MustCompile("^test-group$")),
				),
			},
		},
	})
}
