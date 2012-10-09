// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"github.com/jteeuwen/blah/proto"
	"log"
)

// Bind binds protocol message handlers.
func Bind(c *proto.Client) {
	c.Bind(proto.PIDUnknown, onAny)
	c.Bind(proto.PIDPing, onPing)
	c.Bind(proto.PIDVersion, onVersion)
}

// onAny is a catch-all handler for all incoming messages.
// It is used to write incoming message to a log.
func onAny(c *proto.Client, m *proto.Message) {
	log.Printf("> [%s] %s", m.Type, m.Data)
}

// onPing handles PING messages.
func onPing(c *proto.Client, m *proto.Message) {
	c.Pong(m.Data)
}

// onVersion handles VERSION requests.
func onVersion(c *proto.Client, m *proto.Message) {
	c.Privmsg(m.Receiver, "%s %d.%d.%s",
		AppName, AppVersionMajor, AppVersionMinor, AppVersionRev)
}
