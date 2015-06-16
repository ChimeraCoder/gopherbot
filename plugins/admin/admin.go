// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package admin

import (
	"github.com/chimeracoder/gopherbot/cmd"
	"github.com/chimeracoder/gopherbot/irc"
	"github.com/chimeracoder/gopherbot/plugin"
	"github.com/chimeracoder/gopherbot/proto"
)

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "admin")
	return p
}

func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	comm := new(cmd.Command)
	comm.Name = "quit"
	comm.Description = "Unconditionally quit the bot program"
	comm.Restricted = true
	comm.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		c.Quit("")
	}
	cmd.Register(comm)

	comm = new(cmd.Command)
	comm.Name = "join"
	comm.Description = "Join the given channel"
	comm.Restricted = true
	comm.Params = []cmd.Param{
		{Name: "channel", Optional: false, Pattern: cmd.RegChannel},
		{Name: "key", Optional: true, Pattern: cmd.RegAny},
		{Name: "chanservpass", Optional: true, Pattern: cmd.RegAny},
	}
	comm.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		var ch irc.Channel
		ch.Name = cmd.Params[0].Value

		if len(cmd.Params) > 1 {
			ch.Key = cmd.Params[1].Value
		}

		if len(cmd.Params) > 2 {
			ch.ChanservPassword = cmd.Params[2].Value
		}

		c.Join(&ch)
	}
	cmd.Register(comm)

	comm = new(cmd.Command)
	comm.Name = "leave"
	comm.Description = "Leave the given channel"
	comm.Restricted = true
	comm.Params = []cmd.Param{
		{Name: "channel", Optional: true, Pattern: cmd.RegChannel},
	}
	comm.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		var ch irc.Channel

		if len(cmd.Params) > 0 {
			ch.Name = cmd.Params[0].Value
		} else {
			if !m.FromChannel() {
				return
			}
			ch.Name = m.Receiver
		}

		c.Part(&ch)
	}
	cmd.Register(comm)

	return
}
