resource "jira_group" "foo" {
  name = "test-group"
}

resource "jira_user" "foo" {
  email_address = "foo@joob.dev"
  display_name  = "foo"
}

resource "jira_group_membership" "foo" {
  group_name = jira_group.foo.name
  account_id = jira_user.foo.id
}
