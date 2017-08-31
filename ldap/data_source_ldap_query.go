package ldap

import (
	"fmt"

	ldapapi "gopkg.in/ldap.v2"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceQuery() *schema.Resource {

	return &schema.Resource{

		Read: dataSourceQueryRead,

		Schema: map[string]*schema.Schema{

			"base_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"index_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"results": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"results_attr": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceQueryRead(d *schema.ResourceData, meta interface{}) (err error) {

	var (
		conn *ldapapi.Conn

		searchRequest *ldapapi.SearchRequest
		searchResult  *ldapapi.SearchResult

		keyValue string
	)

	c := meta.(*client)
	if c == nil {
		return fmt.Errorf("ldap client is nil")
	}
	if conn, err = c.connect(); err != nil {
		return
	}
	defer conn.Close()

	baseDN := d.Get("base_dn").(string)
	filter := d.Get("filter").(string)
	attributes := d.Get("attributes").([]interface{})
	indexAttribute := d.Get("index_attribute").(string)

	ldapAttrs := []string{}
	for _, v := range attributes {
		ldapAttrs = append(ldapAttrs, v.(string))
	}

	searchRequest = ldapapi.NewSearchRequest(baseDN,
		ldapapi.ScopeWholeSubtree, ldapapi.DerefAlways, 0, 0, false,
		filter, ldapAttrs, nil)

	if searchResult, err = conn.Search(searchRequest); err != nil {
		return
	}

	id := uuid.New().String()
	results := []interface{}{}
	resultsAttr := make(map[string]interface{})

	for _, entry := range searchResult.Entries {

		c.logDebug("LDAP Search result entry: %# v", entry)

		result := make(map[string]string)
		keyValue = ""

		for _, ldapAttr := range ldapAttrs {

			attrValue := entry.GetAttributeValue(ldapAttr)
			result[ldapAttr] = string(attrValue)
			if ldapAttr == indexAttribute {
				keyValue = attrValue
			}
		}
		if len(keyValue) > 0 {
			for k, v := range result {
				resultsAttr[keyValue+"/"+k] = v
			}
			results = append(results, keyValue)
		}
	}

	d.SetId(id)
	d.Set("results", results)
	d.Set("results_attr", resultsAttr)
	return
}
