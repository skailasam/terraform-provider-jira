package jira

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ctreminiom/go-atlassian/jira"
	atlassian "github.com/ctreminiom/go-atlassian/jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("JIRA_URL", nil),
					Description: "Base url of the JIRA instance.",
				},
				"user": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("JIRA_USER", nil),
					Description: "User to be used",
				},
				"password": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("JIRA_PASSWORD", nil),
					Description: "Password/API Key of the user",
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"jira_user": dataSourceUser(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"jira_group":            resourceGroup(),
				"jira_group_membership": resourceGroupMembership(),
				"jira_issue_type":       resourceIssueType(),
				"jira_project":          resourceProject(),
				"jira_project_category": resourceProjectCategory(),
				"jira_role":             resourceRole(),
				"jira_user":             resourceUser(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type client struct {
	*jira.Client
	user     string
	password string
}

func (c *client) Request(method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.Site.String(), endpoint),
		body,
	)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.user, c.password)
	return c.HTTP.Do(req)
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		c, err := atlassian.New(nil, d.Get("url").(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}
		wrapper := &client{Client: c}
		wrapper.user = d.Get("user").(string)
		wrapper.password = d.Get("password").(string)
		wrapper.Auth.SetBasicAuth(wrapper.user, wrapper.password)
		return wrapper, nil
	}
}
