LDAP Terraform Provider [![Build Status](https://travis-ci.org/mevansam/terraform-provider-ldap.svg?branch=master)](https://travis-ci.org/mevansam/terraform-provider-ldap)
================================

Overview
--------

This Terraform provider plugin interact with LDAP declaratively using [HCL](https://github.com/hashicorp/hcl).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-ldap`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-ldap
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-ldap
$ make build
```

Using the provider
------------------

Download the release binary and copy it to the `$HOME/terraform.d/plugins/<os>_<arch>/` folder. For example `/home/youruser/terraform.d/plugins/linux_amd64` for a Linux environment or `/Users/youruser/terraform.d/plugins/darwin_amd64` for a MacOS environment.

Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

Clone this repository to `GOPATH/src/github.com/terraform-providers/terraform-provider-ldap` as its packaging structure 
has been defined such that it will be compatible with the Terraform provider plugin framwork in 0.10.x.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-ldap
...
```


Testing the Provider
--------------------

In order to run the tests locally an LDAP server needs to be run locally. You can launch this server in Docker with the test directory data using the provided script "`scripts/ldap-up.sh`". Once the LDAP server is running locally it can be accessed via the [phpLDAPadmin](http://phpldapadmin.sourceforge.net/wiki/index.php/Main_Page) application launched alongside the LDAP server via this [link](https://localhost:6443/). Once the instance is running you will need to export the following environment variables.

```
export LDAP_HOST=<local host IP>
export LDAP_BIND_DN=cn=admin,dc=example,dc=org
export LDAP_BIND_PASSWORD=admin
```

Export the following environment variable to enable debug log for the provider.

```
export LDAP_DEBUG=true
```

To launch the LDAP server in docker:

```
cd <repository root>
scripts/ldap-up.sh
```

To run the provider unit tests:

```
cd <repository root>/ldap
TF_ACC=1 go test -v -timeout 120m .
```

>> Acceptance tests are run against a LDAP instance in AWS before a release is created. Any other testing should be done using the local LDAP instance. 

```
$ make testacc
```

Terraform Links
---------------

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)
