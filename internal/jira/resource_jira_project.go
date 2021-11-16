package jira

import (
	"context"

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
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// client := meta.(*jira.Client)
	// client.Project.Create(context.Background(), &jira.ProjectPayloadScheme{
	// 	NotificationScheme  int    `json:"notificationScheme"`
	// 	Description: d.Get("description").(string),
	// 	LeadAccountID: d.Get("lead").(string),
	// 	URL: d.Get("url").(string),
	// 	ProjectTemplateKey: d.Get("template").(string),
	// 	AvatarID            int    `json:"avatarId"`
	// 	IssueSecurityScheme int    `json:"issueSecurityScheme"`
	// 	Name                string `json:"name"`
	// 	PermissionScheme    int    `json:"permissionScheme"`
	// 	AssigneeType        string `json:"assigneeType"`
	// 	ProjectTypeKey      string `json:"projectTypeKey"`
	// 	Key                 string `json:"key"`
	// 	CategoryID          int    `json:"categoryId"`
	// })
	// var cl *confluence.Client
	// cl.Space.Create(context.Background(), &confluence.CreateSpaceScheme{
	// 	Key:  "ADN",
	// 	Name: "Airwallex",
	// 	Description: &confluence.CreateSpaceDescriptionScheme{
	// 		Plain: &confluence.CreateSpaceDescriptionPlainScheme{
	// 			Value:          "",
	// 			Representation: "",
	// 		},
	// 	},
	// 	AnonymousAccess:  false,
	// 	UnlicensedAccess: false,
	// }, false)
	return nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_ = meta.(*jira.Client)
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_ = meta.(*jira.Client)
	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_ = meta.(*jira.Client)
	return nil
}
