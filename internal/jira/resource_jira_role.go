package jira

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceRoleCreate,
		ReadContext:   resourceRoleRead,
		DeleteContext: resourceRoleDelete,

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
				ForceNew:    true,
			},
			"name": {
				Description: "Name",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceRoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	res, body, err := client.Project.Role.Create(context.Background(), &jira.ProjectRolePayloadScheme{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.Errorf("%+v", body.Bytes.String())
	}
	d.SetId(strconv.Itoa(res.ID))
	return resourceRoleRead(ctx, d, meta)
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	roles, _, err := client.Project.Role.Global(context.Background())
	if err != nil {
		return diag.FromErr(err)
	}
	var role *jira.ProjectRoleScheme
	for _, r := range roles {
		if r.ID == id {
			role = r
			break
		}
	}
	if role == nil {
		return diag.Errorf("role not found: %d", id)
	}
	d.Set("name", role.Name)
	d.Set("description", role.Description)
	d.Set("id", strconv.Itoa(role.ID))
	return nil
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)

	buf := bytes.NewReader([]byte{})
	res, err := client.Request(
		http.MethodDelete,
		fmt.Sprintf("rest/api/3/role/%s", d.Id()),
		buf,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != http.StatusNoContent {
		return diag.Errorf("Invalid response code: %d", res.StatusCode)
	}
	return nil
}
