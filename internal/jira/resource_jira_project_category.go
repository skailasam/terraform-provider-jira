package jira

import (
	"context"
	"strconv"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProjectCategory() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceProjectCategoryCreate,
		ReadContext:   resourceProjectCategoryRead,
		UpdateContext: resourceProjectCategoryUpdate,
		DeleteContext: resourceProjectCategoryDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Project category name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Project category name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Project category name.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceProjectCategoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	res, _, err := client.Project.Category.Create(context.Background(), &jira.ProjectCategoryPayloadScheme{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(res.ID)
	return resourceProjectCategoryRead(ctx, d, meta)
}

func resourceProjectCategoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	res, _, err := client.Project.Category.Get(context.Background(), id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", res.Name)
	d.Set("description", res.Description)
	d.Set("id", res.ID)
	return nil
}

func resourceProjectCategoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_, _, err = client.Project.Category.Update(context.Background(), id, &jira.ProjectCategoryPayloadScheme{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceProjectCategoryRead(ctx, d, meta)
}

func resourceProjectCategoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*jira.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = client.Project.Category.Delete(context.Background(), id)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
