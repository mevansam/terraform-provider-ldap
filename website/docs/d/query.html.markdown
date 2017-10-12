---
layout: "ldap"
page_title: "LDAP: ldap_query"
sidebar_current: "docs-ldap-datasource-query"
description: |-
  LDAP query data source.
---

# ldap\_query

Data source for querying LDAP for one or more objects.

## Example Usage

The following example return the email address of all developers in a particular group named "org1".

```
data "ldap_query" "org1" {

  base_dn = "dc=acme,dc=com"
  filter = "(&(objectClass=inetOrgPerson)(memberOf=cn=developers,ou=org1,ou=pcf,dc=example,dc=org))(mail=callison@example.org))"

  attributes = [ "mail", "givenName", "sn" ]
  }

  index_attribute = "mail"
}
```

This is equivalent to the following query using the `ldapsearch` cli.

```
ldapsearch -x -H ldap://myldapserver:389 \
  -D "<bind DN>" -w "<bind password>" \
  -b "dc=example,dc=org" \
  "(&(objectClass=inetOrgPerson)(memberOf=cn=developers,ou=org1,ou=pcf,dc=example,dc=org)(mail=callison@example.org))"
``` 

## Argument Reference

* `base_dn` - (Required, String) The base DN for the query
* `filter` - (Required, String) The LDAP search query filter.

The following arguments declare how the results should be exported so they can be referenced via interpolation.

* `attributes` - (Required, List) The list of the LDAP attributes to be retrieved. 
* `index_attribute` - (Required, String) The LDAP attribute to use to populate the `results` attribute with. The value of this attribute can be used as the key to lookup a LDAP query result record and its attributes.

## Attributes Reference

The following attributes are exported:

* `results` - A list of the values of the `index_attributes` for all entries returned by the query.
* `results_attr` - A map of the LDAP results keyed by the attribute name identified by `key_attribute/<attribute name>`.
