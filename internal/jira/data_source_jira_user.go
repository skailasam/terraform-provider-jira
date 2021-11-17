package jira

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "Search for a user by email and/or username.",
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)

	v, ok := d.GetOk("email_address")
	if !ok {
		return diag.Errorf("Expected an email_address to search for")
	}

	users, _, err := client.User.Search.Do(context.Background(), "", v.(string), 0, 2)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(users) != 1 {
		return diag.Errorf("Expected a single result, received %d results.", len(users))
	}

	var user = users[0]

	d.SetId(user.AccountID)
	d.Set("id", user.AccountID)
	d.Set("display_name", user.DisplayName)
	d.Set("email_address", user.EmailAddress)
	return nil
}
