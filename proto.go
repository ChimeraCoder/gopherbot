// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"github.com/jteeuwen/ircb/cmd"
	"github.com/jteeuwen/ircb/proto"
	"log"
)

// bind binds protocol message handlers.
func bind(c *proto.Client) {
	c.Bind(proto.Unknown, onAny)
	c.Bind(proto.CmdPing, onPing)
	c.Bind(proto.EndOfMOTD, onJoinChannels)
	c.Bind(proto.ErrNoMOTD, onJoinChannels)
	c.Bind(proto.ErrNicknameInUse, onNickInUse)
	c.Bind(proto.CmdPrivMsg, onPrivMsg)
}

// onAny is a catch-all handler for all incoming messages.
// It is used to write incoming messages to a log.
func onAny(c *proto.Client, m *proto.Message) {
	log.Printf("> [%03d] [%s] %s", m.Command, m.Receiver, m.Data)
}

// onPing handles PING messages.
func onPing(c *proto.Client, m *proto.Message) {
	c.Pong(m.Data)
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

// onPrivMsg handles private messages directed at us.
// We want to know if it concerns a CTCP request, a bot command
// or just random talk.
func onPrivMsg(c *proto.Client, m *proto.Message) {
	switch {
	case cmd.Parse(config.CommandPrefix, c, m):
	case ctcpVersion(c, m):
	case ctcpPing(c, m):
	}
}
