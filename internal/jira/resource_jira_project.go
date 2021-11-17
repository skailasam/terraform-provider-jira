package jira

import (
	"context"
	"errors"
	"strconv"

	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"key": {
				Description: "Project key.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"lead": {
				Description: "Lead account ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Project name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Project description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"category_id": {
				Description: "Category ID",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"avatar_id": {
				Description: "Avatar ID",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10308,
			},
			"notification_scheme": {
				Description: "Notification Scheme",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"issue_security_scheme": {
				Description: "Avatar ID",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"permission_scheme": {
				Description: "Avatar ID",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"assignee_type": {
				Description: "Assignee type",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "UNASSIGNED",
			},
			"template_key": {
				Description: "Template key.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban",
			},
			"type_key": {
				Description: "Type key",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "software",
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)

	data := &jira.ProjectPayloadScheme{
		Key:                 d.Get("key").(string),
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		LeadAccountID:       d.Get("lead").(string),
		AvatarID:            d.Get("avatar_id").(int),
		AssigneeType:        d.Get("assignee_type").(string),
		ProjectTemplateKey:  d.Get("template_key").(string),
		ProjectTypeKey:      d.Get("type_key").(string),
		CategoryID:          d.Get("category_id").(int),
		NotificationScheme:  d.Get("notification_scheme").(int),
		IssueSecurityScheme: d.Get("issue_security_scheme").(int),
		PermissionScheme:    d.Get("permission_scheme").(int),
	}

	project, res, err := client.Project.Create(context.Background(), data)
	if err != nil {
		return diag.FromErr(errors.New(res.Bytes.String()))
	}
	d.SetId(strconv.Itoa(project.ID))
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	project, _, err := client.Project.Get(context.Background(), d.Id(), []string{})
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("id", project.ID)
	d.Set("key", project.Key)
	d.Set("name", project.Name)
	d.Set("description", project.Description)
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	_, _, err := client.Project.Update(context.Background(), d.Id(), &jira.ProjectUpdateScheme{
		Key:         d.Get("key").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client)
	_, err := client.Project.Delete(context.Background(), d.Id(), false)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
