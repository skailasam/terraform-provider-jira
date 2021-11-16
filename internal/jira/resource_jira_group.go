package jira

import (
	"context"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		DeleteContext: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Group name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"members": {
				Description: "Member list",
				Computed:    true,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Description: "Account ID",
					Type:        schema.TypeString,
				},
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	name := d.Get("name").(string)
	_, _, err := client.Group.Create(context.Background(), name)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceGroupRead(ctx, d, meta)
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	name := d.Get("name").(string)
	group, _, err := client.Group.Members(context.Background(), name, false, 0, 1000)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)
	var members []string
	for _, member := range group.Values {
		members = append(members, member.AccountID)
	}
	d.Set("members", members)
	return nil
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	_, err := client.Group.Delete(context.Background(), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
