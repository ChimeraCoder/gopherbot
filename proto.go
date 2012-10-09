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
	c.Bind(proto.PIDEndOfMOTD, onJoinChannels)
	c.Bind(proto.PIDNoMOTD, onJoinChannels)
	c.Bind(proto.PIDNickInUse, onNickInUse)
}

// onAny is a catch-all handler for all incoming messages.
// It is used to write incoming messages to a log.
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

// onJoinChannels is used to complete the login procedure.
// We have just received the server's MOTD and now is a good time to
// start joining channels.
func onJoinChannels(c *proto.Client, m *proto.Message) {
	c.Join(config.Channels)
}

// onNickInUse is called whenever we receive a notification that our
// nickname is already in use. We will attempt to re-acquire it by 
// identifying with our password. Otherwise we will pick a new name.
func onNickInUse(c *proto.Client, m *proto.Message) {
	if len(config.NickservPassword) > 0 {
		c.Recover(config.Nickname, config.NickservPassword)
		return
	}

	config.SetNickname(config.Nickname + "_")
	c.Nick(config.Nickname, "")
}
