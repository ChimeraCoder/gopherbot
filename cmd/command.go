// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"github.com/chimeracoder/gopherbot/proto"
	"strings"
)

var (
	// List of registered commands.
	commands []*Command

	// User whitelist
	whitelist []string
)

// Register registers the given command name and constructor.
// Modules should call this during initialization to register their
// commands with the bot.
func Register(c *Command) { commands = append(commands, c) }

// SetWhitelist sets the list of user hostmasks. These users are allowed to
// execute restricted commands.
func SetWhitelist(list []string) { whitelist = list }

// findCommand finds the first command instance for the given name.
func findCommand(name string) *Command {
	for _, c := range commands {
		if strings.EqualFold(name, c.Name) {
			return c.Copy()
		}
	}

	return nil
}

// isWhitelisted returns true if the given name is in the user whitelist.
func isWhitelisted(name string) bool {
	for _, mask := range whitelist {
		if strings.EqualFold(name, mask) {
			return true
		}
	}

	return false
}

// CommandFunc represents a command constructor.
type CommandFunc func() *Command

// ExecuteFunc represents a command execution handler.
// These are executed in a separate goroutine.
type ExecuteFunc func(*Command, *proto.Client, *proto.Message)

// Command represents a single bot command.
type Command struct {
	Name        string      // Command name.
	Description string      // Command description.
	Data        string      // Original parameter data as a single string.
	Params      []Param     // Command parameters.
	Execute     ExecuteFunc // Execution handler for the command.
	Restricted  bool        // Command is restricted to admin users.
}

// Copy returns a deep copy of the current command.
func (c *Command) Copy() *Command {
	nc := new(Command)
	nc.Name = c.Name
	nc.Description = c.Description
	nc.Execute = c.Execute
	nc.Restricted = c.Restricted
	nc.Params = make([]Param, len(c.Params))

	for i := range c.Params {
		nc.Params[i] = *c.Params[i].Copy()
	}

	return nc
}

// RequiredParamCount counts the number of required parameters.
func (c *Command) RequiredParamCount() int {
	var pc int

	for i := range c.Params {
		if !c.Params[i].Optional {
			pc++
		}
	}

	return pc
}
