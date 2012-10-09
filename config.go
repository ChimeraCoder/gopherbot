// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"github.com/jteeuwen/blah/irc"
	"github.com/jteeuwen/ini"
)

// Config holds bot configuration data.
type Config struct {
	Address          string
	SSLKey           string
	SSLCert          string
	Nickname         string
	ServerPassword   string
	OperPassword     string
	NickservPassword string
	QuitMessage      string
	Channels         []*irc.Channel
}

// Load loads configuration data from the given ini file.
func (c *Config) Load(file string) (err error) {
	ini := ini.New()
	err = ini.Load(file)

	if err != nil {
		return
	}

	s := ini.Section("net")
	c.Address = fmt.Sprintf("%s:%d", s.S("host", ""), s.U32("port", 0))
	c.SSLKey = s.S("ssl-key", "")
	c.SSLCert = s.S("ssl-cert", "")

	s = ini.Section("account")
	c.Nickname = s.S("nickname", "")
	c.ServerPassword = s.S("server-password", "")
	c.OperPassword = s.S("oper-password", "")
	c.NickservPassword = s.S("nickserv-password", "")
	c.QuitMessage = s.S("quit-message", "")
	return
}
