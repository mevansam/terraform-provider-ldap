---
layout: "ldap"
page_title: "Provider: LDAP"
sidebar_current: "docs-ldap-index"
description: |-
  An LDAP provider to manage an LDAP backend.
---

# LDAP Provider

The LDAP provider can be used to query and manage and LDAP backend.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Set the variable values in *.tfvars file
# or using -var="api_url=..." CLI option

variable "ldap_host" {}
variable "ldap_bind_dn" {}
variable "ldap_bind_password" {}

# Configure the Cloud Foundry LDAP Provider

provider "ldap" {
    
    host = ${var.ldap_host}"
    port = 389

    bind_dn = "${var.ldap_bind_dn}"
    bind_password = "${var.ldap_bind_password}"

    skip_ssl_validation = true
}
```

## Argument Reference

The following arguments are supported:

* `host` - (Required, String) The DNS name or IP of the LDAP backend. This can also be specified with the `LDAP_HOST` shell environment variable.

* `port` - (Optional, Number) The port on which the LDAP service is listening. The default is 389 if `use_tls` is false and 636 if it is true. This can also be specified with the `LDAP_PORT` shell environment variable.

* `use_tls` - (Optional, Boolean) Whether to use TLS when connecting to the LDAP service. The default is "false". This can also be specified with the `LDAP_USE_TLS` shell environment variable.

* `tls_skip_verify` - (Optional) Skip TLS verification. Defaults to "false". This can also be specified with the `LDAP_TLS_SKIP_VERIFY` shell environment variable.

* `bind_dn` - (Required) The BIND DN to use when binding to the LDPA backend. This can also be specified
  with the `LDAP_BIND_DN` shell environment variable.

* `bind_password` - (Required) The password of the BIND DN. This can also be specified
  with the `LDAP_BIND_PASSWORD` shell environment variable.
