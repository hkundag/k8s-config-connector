package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
