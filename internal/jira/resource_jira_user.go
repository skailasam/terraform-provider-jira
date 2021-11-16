package jira

import (
	"context"

	"errors"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"email_address": {
				Description: "E-mail address",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"display_name": {
				Description: "Display name",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: func(interface{}, string) ([]string, []error) {
					return []string{}, nil
				},
			},
			"id": {
				Description: "Account ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	user, res, err := client.User.Create(context.Background(), &jira.UserPayloadScheme{
		EmailAddress: d.Get("email_address").(string),
		DisplayName:  d.Get("display_name").(string),
		Notification: false,
	})
	if err != nil {
		return diag.FromErr(errors.New(res.Bytes.String()))
	}
	d.SetId(user.AccountID)
	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	user, _, err := client.User.Get(context.Background(), d.Id(), []string{})
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("id", user.AccountID)
	d.Set("display_name", user.DisplayName)
	d.Set("email_address", user.EmailAddress)
	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	_, err := client.User.Delete(context.Background(), d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
