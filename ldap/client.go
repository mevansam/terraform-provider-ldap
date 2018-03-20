package ldap

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"time"

	"github.com/kr/pretty"

	ldapapi "gopkg.in/ldap.v2"
)

// client -
type client struct {
	host          string
	port          int
	useTLS        bool
	tlsSkipVerify bool
	bindDN        string
	bindPasswod   string

	debug bool
}

// connect -
func (c *client) connect() (conn *ldapapi.Conn, err error) {

	if c.useTLS {
		if c.port == -1 {
			c.port = 636
		}

		conn, err = ldapapi.DialTLS("tcp",
			fmt.Sprintf("%s:%d", c.host, c.port),
			&tls.Config{
				InsecureSkipVerify: c.tlsSkipVerify,
				ServerName:         c.host,
			})

	} else {
		if c.port == -1 {
			c.port = 389
		}

		conn, err = ldapapi.Dial("tcp",
			fmt.Sprintf("%s:%d", c.host, c.port))
	}
	if err != nil {
		return
	}

	if err = conn.Bind(c.bindDN, c.bindPasswod); err != nil {
		return
	}
	return
}

// logDebug -
func (c *client) logDebug(format string, v ...interface{}) {
	if c.debug {
		vv := []interface{}{}
		for _, o := range v {
			k := reflect.ValueOf(o).Kind()
			if k == reflect.Struct ||
				k == reflect.Interface ||
				k == reflect.Ptr ||
				k == reflect.Slice ||
				k == reflect.Map {
				vv = append(vv, pretty.Formatter(o))
			} else {
				vv = append(vv, o)
			}
		}
		hdr := fmt.Sprintf("[%s] DEBUG:", time.Now().Format(time.RFC3339))
		fmt.Printf(fmt.Sprintf("%s %s\n", hdr, format), vv...)
	}
}
