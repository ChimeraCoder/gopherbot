// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package dict

import (
	"bytes"
	"github.com/jteeuwen/ircb/cmd"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"log"
)

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "dict")
	return p
}

func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	w := new(cmd.Command)
	w.Name = "define"
	w.Description = "Fetch the definition for the given term"
	w.Restricted = false
	w.Params = []cmd.Param{
		{Name: "term", Description: "Word to find definition for", Pattern: cmd.RegAny},
	}
	w.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		dict, err := Dial("tcp", "dict.org:2628")
		if err != nil {
			log.Printf("[dict] %s", err)
			c.PrivMsg(m.Receiver, "%s, No definition found for '%s'",
				m.SenderName, cmd.Params[0].Value)
			return
		}

		def, err := dict.Define("wn", cmd.Params[0].Value)
		if err != nil {
			log.Printf("[dict] %s", err)
			c.PrivMsg(m.Receiver, "%s, No definition found for '%s'",
				m.SenderName, cmd.Params[0].Value)
			return
		}

		if len(def) == 0 {
			c.PrivMsg(m.Receiver, "%s, No definition found for '%s'",
				m.SenderName, cmd.Params[0].Value)
			return
		}

		space := []byte{' '}
		mspace := []byte{' ', ' '}
		line := bytes.Replace(def[0].Text, []byte{'\n'}, space, -1)

		// Strip all multi-space indents.
		for bytes.Index(line, mspace) > -1 {
			line = bytes.Replace(line, mspace, space, -1)
		}

		c.PrivMsg(m.Receiver, "%s: %s", m.SenderName, line)
	}

	cmd.Register(w)

	return
}
