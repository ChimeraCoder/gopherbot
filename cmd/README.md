## cmd

This package holds bot command parsing and execution code.
It is used by registering a command and handler with this package, before
initializing the connection. Once the connection is active, the command
package must be invoked for every `PRIVMSG` request:

	client.Bind(proto.CmdPrivMsg, onPrivMsg)
	
	...
	
	func onPrivMsg(c *proto.Client, m *proto.Message) {
		...
		cmd.Parse(commandPrefix, c, m)
		...
	}

The incoming message will be processed and matched against any registered
commands. A command's execution handler will be invoked if, and only if the
supplied message data matches all parameters, and the user sendering the
request has permission to execute the command.

If the user omits required arguments, or the supplied arguments do not match
the format we expect them to have, the bot will automatically send an
appropriate error response to the user and the command handler is not invoked.


### Examples

Register a command without any parameters:

	cmd.Register("help", func() *cmd.Command {
		c := new(cmd.Command)
		c.Name = "help"
		c.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
			// Code to handle command execution goes here.
		}
		return c
	})

This can be invoked through IRC by sending `!help`.
Provided `!` is registered as the current command prefix.

Register a command with two parameters:

	cmd.Register("add", func() *cmd.Command {
		c := new(cmd.Command)
		c.Name = "add"
		c.Params = []cmd.Param{
			{Name: "a", Pattern: cmd.RegDecimal},
			{Name: "b", Pattern: cmd.RegDecimal},
		}
		c.Execute = func(cmd *Command, c *proto.Client, m *proto.Message) {
			c.PrivMsg(m.SenderName, "%f", cmd.Params[0].F64(0)+cmd.Params[1].F64(0))
		}
		return c
	})

This can be invoked through IRC by sending `!add 1.23 3.56`.
Provided `!` is registered as the current command prefix.


### Command Parameters

A command which has parameters, is given a reference to the parsed parameters
when its handler is executed. The values for these parameters have already been
validated against a specified regular expression pattern. Once the command
handler is executed, we can therefore be reasonably certain that we received
all the data needed to safely execute the command.

Accessing the parameters is done through the `Command.Params` slice.
Each parameter has convenience methods allowing quick conversion to a given
data type. For example:

		_ = cmd.Params[0].F64(0) + cmd.Params[1].F64(0)

In this code, the two required parameters are referenced as `float64` values.
There are convenience accessors for all the basic Go data types:

	Param.Bool(bool) bool
	Param.I8(int8) int8
	Param.I16(int16) int16
	Param.I32(int32) int32
	Param.I64(int64) int64
	Param.U8(uint8) uint8
	Param.U16(uint16) uint16
	Param.U32(uint32) uint32
	Param.U64(uint64) uint64
	Param.F32(float32) float32
	Param.F64(float64) float64

The one parameter for these is the default value returned when the data type
conversion fails. To access the original string value, simply use the
`Param.Value` field directly.


### Usage

    go get github.com/jteeuwen/ircb/cmd


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

