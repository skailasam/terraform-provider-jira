package jira

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceGroupMembershipCreate,
		ReadContext:   resourceGroupMembershipRead,
		DeleteContext: resourceGroupMembershipDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Description: "Group name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Description: "Account ID",
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceGroupMembershipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	n := d.Get("group_name").(string)
	a := d.Get("account_id").(string)
	_, _, err := client.Group.Add(context.Background(), n, a)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s:%s", n, a))
	return nil
}

func resourceGroupMembershipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	name := d.Get("group_name").(string)
	members, _, err := client.Group.Members(context.Background(), name, true, 0, 1000)
	if err != nil {
		return diag.FromErr(err)
	}
	id := d.Get("account_id").(string)
	var found bool
	for _, member := range members.Values {
		if member.AccountID == id {
			found = true
			break
		}
	}
	if !found {
		return diag.Errorf("invalid membership for group %s and member %s", name, id)
	}
	return nil
}

func resourceGroupMembershipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	n := d.Get("group_name").(string)
	a := d.Get("account_id").(string)
	_, err := client.Group.Remove(context.Background(), n, a)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
