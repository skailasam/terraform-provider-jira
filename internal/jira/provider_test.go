package jira

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"jira": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("JIRA_URL"); v == "" {
		t.Fatal("JIRA_URL must be set for acceptance tests")
	}

	if v := os.Getenv("JIRA_USER"); v == "" {
		t.Fatal("JIRA_USER must be set for acceptance tests")
	}

	if v := os.Getenv("JIRA_PASSWORD"); v == "" {
		t.Fatal("JIRA_PASSWORD must be set for acceptance tests")
	}
}
