package jira

import (
	"context"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIssueType() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceIssueTypeCreate,
		ReadContext:   resourceIssueTypeRead,
		UpdateContext: resourceIssueTypeUpdate,
		DeleteContext: resourceIssueTypeDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hierarchy_level": {
				Description: "Hierarchy level",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Computed:    false,
			},
		},
	}
}

func resourceIssueTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	res, _, err := client.Issue.Type.Create(context.Background(), &jira.IssueTypePayloadScheme{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		HierarchyLevel: d.Get("hierarchy_level").(int),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(res.ID)
	return resourceIssueTypeRead(ctx, d, meta)
}

func resourceIssueTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	res, _, err := client.Issue.Type.Get(context.Background(), d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", res.Name)
	d.Set("description", res.Description)
	d.Set("hierarchy_level", res.HierarchyLevel)
	d.Set("id", res.ID)
	return nil
}

func resourceIssueTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	_, _, err := client.Issue.Type.Update(context.Background(), d.Id(), &jira.IssueTypePayloadScheme{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		HierarchyLevel: d.Get("hierarchy_level").(int),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIssueTypeRead(ctx, d, meta)
}

func resourceIssueTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	_, err := client.Issue.Type.Delete(context.Background(), d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
