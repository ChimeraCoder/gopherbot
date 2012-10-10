// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"github.com/jteeuwen/ini"
	"github.com/jteeuwen/ircb/irc"
	"strings"
	"sync/atomic"
	"unsafe"
)

// Global bot configuration settings.
var config *Config

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

// SetNickname atomically sets the new nickname.
// This is used in response to proto.PIDNickInUse messages.
func (c *Config) SetNickname(nickname string) {
	new := *c
	new.Nickname = nickname

	for {
		old := (*unsafe.Pointer)(unsafe.Pointer(&config))

		if atomic.CompareAndSwapPointer(old, *old, unsafe.Pointer(&new)) {
			return
		}
	}

	panic("Unable to change nickname")
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
	c.SSLKey = s.S("x509-key", "")
	c.SSLCert = s.S("x509-cert", "")

	chans := s.List("channels")
	c.Channels = make([]*irc.Channel, len(chans))

	// Parse channel definitions. A single channel comes as a string like:
	//
	//    <name>,<key>,<chanservpassword>
	//
	// The name is the only required value.
	for i, line := range chans {
		elements := strings.Split(line, ",")

		for k := range elements {
			elements[k] = strings.TrimSpace(elements[k])
		}

		if len(elements) == 0 || len(elements[0]) == 0 {
			continue
		}

		var ch irc.Channel
		ch.Name = elements[0]

		if len(elements) > 1 {
			ch.Key = elements[1]
		}

		if len(elements) > 2 {
			ch.ChanservPassword = elements[2]
		}

		c.Channels[i] = &ch
	}

	s = ini.Section("account")
	c.Nickname = s.S("nickname", "")
	c.ServerPassword = s.S("server-password", "")
	c.OperPassword = s.S("oper-password", "")
	c.NickservPassword = s.S("nickserv-password", "")
	c.QuitMessage = s.S("quit-message", "")
	return
}
