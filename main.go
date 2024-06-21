package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/glovecchi0/terraform-provider-neuvector/neuvector"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: neuvector.Provider,
    })
}
