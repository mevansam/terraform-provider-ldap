package ldap

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ldap": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {

	if !testAccEnvironmentSet() {
		t.Fatal("Acceptance environment has not been set.")
	}
}

func testAccEnvironmentSet() bool {

	host := os.Getenv("LDAP_HOST")
	bindDN := os.Getenv("LDAP_BIND_DN")
	bindPasswod := os.Getenv("LDAP_BIND_PASSWORD")

	if len(host) == 0 ||
		len(bindDN) == 0 ||
		len(bindPasswod) == 0 {

		fmt.Println("LDAP_HOST, LDAP_BIND_DN and LDAP_BIND_PASSWORD " +
			"must be set for acceptance tests to work.")
		return false
	}
	return true
}
