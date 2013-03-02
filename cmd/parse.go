// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"bytes"
	"github.com/jteeuwen/ircb/proto"
	"strings"
	"text/scanner"
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
	for i := 0; i < pc && i < lp; i++ {
		cmd.Params[i].Value = params[i]

		if !cmd.Params[i].Valid() {
			c.PrivMsg(m.SenderName, "Invalid parameter value %q for command %q",
				params[i], name)
			return false
		}
	}

	// Execute the command.
	if cmd.Execute == nil {
		c.PrivMsg(m.SenderName, "Command %q is not implemented", name)
		return false
	}

	go cmd.Execute(cmd, c, m)
	return true
}

// parseCommand reads command name and arguments from the given input.
func parseCommand(data string) (string, []string) {
	var scan scanner.Scanner
	var list []string

	scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.SkipComments |
		scanner.ScanRawStrings | scanner.ScanChars
	scan.Init(bytes.NewBufferString(data))

	tok := scan.Scan()
	for tok != scanner.EOF {
		if scan.ErrorCount > 0 {
			break
		}

		list = append(list, scan.TokenText())
		tok = scan.Scan()
	}

	if len(list) == 0 {
		return "", nil
	}

	return strings.ToLower(list[0]), list[1:]
}
