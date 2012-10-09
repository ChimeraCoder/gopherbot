// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"github.com/jteeuwen/ircb/irc"
	"github.com/jteeuwen/ini"
	"strings"
	"sync/atomic"
	"unsafe"
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

// SetNickname atomically sets the new nickname.
// This is used in response to proto.PIDNickInUse messages.
func (c *Config) SetNickname(nickname string) {
	new := unsafe.Pointer(&nickname)

	// FIXME(jimt): Find out how this is supposed to work.
	for i := 0; i < 5; i++ {
		old := unsafe.Pointer(&c.Nickname)
		val := (*unsafe.Pointer)(old)

		if atomic.CompareAndSwapPointer(val, old, new) {
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
	c.SSLKey = s.S("ssl-key", "")
	c.SSLCert = s.S("ssl-cert", "")

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
