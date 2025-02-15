package main

import (
	"github.com/rosineygp/terraform-provider-http-client/httpclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return httpclient.Provider()
		},
	})
}
