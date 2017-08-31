package ldap

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const ldapQuerySingleUser = `

data "ldap_query" "user1" {

  base_dn = "dc=example,dc=org"
  filter = "(&(objectClass=inetOrgPerson)(memberOf=cn=developers,ou=org1,ou=pcf,dc=example,dc=org)(mail=callison@example.org))"

  attributes = [ "mail", "givenName", "sn" ]
  index_attribute = "mail"
}
`

const ldapQueryUserGroup = `

data "ldap_query" "org2" {

  base_dn = "dc=example,dc=org"
  filter = "(&(objectClass=inetOrgPerson)(memberOf=cn=developers,ou=org2,ou=pcf,dc=example,dc=org))"

  attributes = [ "uid", "mail" ]
  index_attribute = "uid"
}
`

func TestAccDataSourceLDAPQuerySingleUser_normal(t *testing.T) {

	ref := "data.ldap_query.user1"

	resource.Test(t,
		resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{

				resource.TestStep{
					Config: ldapQuerySingleUser,
					Check: resource.ComposeTestCheckFunc(
						checkDataSourceLdapQuery(ref),
						resource.TestCheckResourceAttr(
							ref, "results.#", "1"),
						resource.TestCheckResourceAttr(
							ref, "results.0", "callison@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.%", "3"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.callison@example.org/mail", "callison@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.callison@example.org/givenName", "Conrad"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.callison@example.org/sn", "Allison"),
					),
				},
			},
		})

}

func TestAccDataSourceLDAPQueryUserGroup_normal(t *testing.T) {

	ref := "data.ldap_query.org2"
	resource.Test(t,
		resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{

				resource.TestStep{
					Config: ldapQueryUserGroup,
					Check: resource.ComposeTestCheckFunc(
						checkDataSourceLdapQuery(ref),
						resource.TestCheckResourceAttr(
							ref, "results.#", "7"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.%", "14"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.jharrington/uid", "jharrington"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.jharrington/mail", "jharrington@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.hquinn/uid", "hquinn"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.hquinn/mail", "hquinn@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.dford/uid", "dford"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.dford/mail", "dford@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.gmckenzie/uid", "gmckenzie"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.gmckenzie/mail", "gmckenzie@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.nmeyer/uid", "nmeyer"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.nmeyer/mail", "nmeyer@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.mwillis/uid", "mwillis"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.mwillis/mail", "mwillis@example.org"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.lconner/uid", "lconner"),
						resource.TestCheckResourceAttr(
							ref, "results_attr.lconner/mail", "lconner@example.org"),
					),
				},
			},
		})
}

func checkDataSourceLdapQuery(resource string) resource.TestCheckFunc {

	return func(s *terraform.State) error {

		c := testAccProvider.Meta().(*client)
		if c == nil {
			return fmt.Errorf("LDAP connection is not valid")
		}

		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("user '%s' not found in terraform state", resource)
		}
		c.logDebug("terraform state for resource '%s': %# v", resource, rs)

		attributes := rs.Primary.Attributes
		c.logDebug("terraform state for resource '%s': %# v", resource, attributes)

		return nil
	}
}
