package neuvector

import (
    "context"

    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "username": {
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
            },
            "password": {
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
                Sensitive: true,
            },
            "url": {
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "neuvector_application": resourceApplication(),
        },
        ConfigureContextFunc: providerConfigure,
    }
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
    username := d.Get("username").(string)
    password := d.Get("password").(string)
    url := d.Get("url").(string)

    config := Config{
        Username: username,
        Password: password,
        URL:      url,
    }

    return &config, nil
}

type Config struct {
    Username string
    Password string
    URL      string
}
