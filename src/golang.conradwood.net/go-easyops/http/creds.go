package http

import (
	"fmt"
)

type creds struct {
	username string
	password string
}

type cred_producer struct {
	host        string
	known_creds []*creds
	used        int
}

func (c *cred_producer) AddUsernamePassword(username, password string) {
	cr := &creds{username: username, password: password}
	c.known_creds = append(c.known_creds, cr)
}
func (c *cred_producer) getNetRC() *creds {
	fmt.Printf("Getting netrc for host \"%s\"\n", c.host)
	cr := &creds{username: "foo", password: "bar"}
	return cr
}
func (c *cred_producer) GetCredentials() *creds {
	if c.used == len(c.known_creds) {
		c.used++
		return c.getNetRC()
	}
	if c.used > len(c.known_creds) {
		return nil
	}
	res := c.known_creds[c.used]
	c.used++
	return res
}
