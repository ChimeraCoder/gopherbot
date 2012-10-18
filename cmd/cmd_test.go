// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"bytes"
	"github.com/jteeuwen/ircb/proto"
	"testing"
)

const Prefix = "?"

func TestHelp(t *testing.T) {
	Register("help", func() *Command {
		c := new(Command)
		c.Name = "help"
		c.Execute = func(cmd *Command, c *proto.Client, m *proto.Message) {

		}
		return c
	})

	var buf bytes.Buffer
	client := proto.NewClient(func(p []byte) error {
		_, err := buf.Write(p)
		return err
	})

	client.Bind(proto.CmdPrivMsg, func(c *proto.Client, m *proto.Message) {
		if !Parse(Prefix, c, m) {
			t.Fatalf("%s", buf.String())
		}
	})

	client.Read(":steve!b@c.com PRIVMSG bob :?help")
}

func TestAdd(t *testing.T) {
	Register("add", func() *Command {
		c := new(Command)
		c.Name = "add"
		c.Params = []Param{
			Param{Name: "a", Pattern: RegDecimal},
			Param{Name: "b", Pattern: RegDecimal},
		}
		c.Execute = func(cmd *Command, c *proto.Client, m *proto.Message) {
			c.PrivMsg(m.SenderName, "%d", cmd.Params[0].F64(0)+cmd.Params[1].F64(0))
		}
		return c
	})

	var buf bytes.Buffer
	client := proto.NewClient(func(p []byte) error {
		_, err := buf.Write(p)
		return err
	})

	client.Bind(proto.CmdPrivMsg, func(c *proto.Client, m *proto.Message) {
		if !Parse(Prefix, c, m) {
			t.Fatalf("%s", buf.String())
		}
	})

	client.Read(":steve!b@c.com PRIVMSG bob :?ADD 1 2")
}
