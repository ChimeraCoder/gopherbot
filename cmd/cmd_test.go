// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"testing"
)

func Test(t *testing.T) {
	Register("help", makeHelp)
}

func makeHelp() *Command {
	c := new(Command)
	c.Name = "help"
	c.Description = "This command displays command help"
	c.Params = nil
	c.Restricted = false
	return c
}
