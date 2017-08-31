package ldap

import (
	"os"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider -
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_HOST", ""),
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_PORT", -1),
			},
			"use_tls": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_USE_TLS", false),
			},
			"tls_skip_verify": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_TLS_SKIP_VERIFY", false),
			},
			"bind_dn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_BIND_DN", ""),
			},
			"bind_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_BIND_PASSWORD", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ldap_query": dataSourceQuery(),
		},

		ResourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	debug, _ := strconv.ParseBool(os.Getenv("LDAP_DEBUG"))

	return &client{
		host:          d.Get("host").(string),
		port:          d.Get("port").(int),
		useTLS:        d.Get("use_tls").(bool),
		tlsSkipVerify: d.Get("tls_skip_verify").(bool),
		bindDN:        d.Get("bind_dn").(string),
		bindPasswod:   d.Get("bind_password").(string),

		debug: debug,
	}, nil
}
