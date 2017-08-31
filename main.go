package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ldap/ldap"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ldap.Provider})

}
