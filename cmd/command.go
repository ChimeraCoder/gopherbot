// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"github.com/jteeuwen/ircb/proto"
	"strings"
)

// CommandFunc represents a command constructor.
type CommandFunc func() *Command

// Command represents a single bot command.
type Command struct {
	Name        string  // Command name.
	Description string  // Command description.
	Params      []Param // Command parameters.
	Restricted  bool    // Command is restricted to admin users.
}

// List of registered command constructors.
var commands = make(map[string]CommandFunc)

// Register registers the given command name and constructor.
// Modules should call this during initialization to register their
// commands with the bot.
func Register(name string, cf CommandFunc) {
	commands[name] = cf
}

// Parse reads incoming message data and tries to parse it into
// a command structure and then execute it.
func Parse(prefix string, c *proto.Client, m *proto.Message) bool {
	prefixlen := len(prefix)

	if prefixlen == 0 || !strings.HasPrefix(m.Data, prefix) {
		return false
	}

	m.Data = m.Data[prefixlen:]
	return true
}
