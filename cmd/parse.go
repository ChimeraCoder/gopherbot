// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"github.com/jteeuwen/ircb/proto"
	"strings"
)

// Parse reads incoming message data and tries to parse it into
// a command structure and then execute it.
func Parse(prefix string, c *proto.Client, m *proto.Message) bool {
	prefixlen := len(prefix)

	if prefixlen == 0 || !strings.HasPrefix(m.Data, prefix) {
		return false
	}

	// Split the data into a name and list of parameters.
	name, params := parseCommand(m.Data[prefixlen:])
	if len(name) == 0 {
		return false
	}

	// Ensure the given command exists.
	cmd := findCommand(name)
	if cmd == nil {
		return false
	}

	// Ensure the current user us allowed to execute the command.
	if cmd.Restricted && !isWhitelisted(m.SenderMask) {
		c.PrivMsg(m.SenderName, "Access to %q denied.", name)
		return false
	}

	// Make sure we received enough parameters.
	pc := cmd.RequiredParamCount()
	lp := len(params)

	if pc > lp {
		c.PrivMsg(m.SenderName, "Missing parameters for command %q", name)
		return false
	}

	// Copy over parameter values and ensure they are of the right format.
	for i := 0; i < lp && i < len(cmd.Params); i++ {
		cmd.Params[i].Value = params[i]

		if !cmd.Params[i].Valid() {
			c.PrivMsg(m.SenderName, "Invalid parameter value %q for command %q",
				params[i], name)
			return false
		}
	}

	// Execute the command.
	if cmd.Execute != nil {
		go cmd.Execute(cmd, c, m)
	}

	return true
}

// parseCommand reads command name and arguments from the given input.
func parseCommand(data string) (string, []string) {
	var list []string
	var quoted bool

loop:
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '"':
			quoted = !quoted
			fallthrough

		case ' ', '\t':
			if quoted {
				break
			}

			v := strings.TrimSpace(data[:i])
			if len(v) == 0 {
				break
			}

			if v[0] == '"' {
				v = v[1:]

				if len(v) == 0 {
					break
				}
			}

			list = append(list, v)
			data = data[i+1:]
			goto loop
		}
	}

	v := strings.TrimSpace(data)
	if len(v) > 0 {
		if v[0] == '"' {
			v = v[1:]
		}

		if len(v) > 0 {
			list = append(list, v)
		}
	}

	if len(list) == 0 {
		return "", nil
	}

	return strings.ToLower(list[0]), list[1:]
}
