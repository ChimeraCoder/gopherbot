// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"github.com/jteeuwen/ircb/proto"
	"strings"
)

func ctcpVersion(c *proto.Client, m *proto.Message) bool {
	if m.Data != "\x01VERSION\x01" {
		return false
	}

	c.PrivMsg(m.Receiver, "%s %d.%d", AppName, AppVersionMajor, AppVersionMinor)
	return true
}

func ctcpPing(c *proto.Client, m *proto.Message) bool {
	if !strings.HasPrefix(m.Data, "\x01PING ") {
		return false
	}

	c.PrivMsg(m.Receiver, "\x01PONG %s\x01", m.Data[6:])
	return true
}
