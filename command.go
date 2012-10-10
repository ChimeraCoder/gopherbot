// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"github.com/jteeuwen/ircb/proto"
	"strings"
)

// readCommand reads incoming message data and tries to parse it into
// a command structure and then execute it.
func readCommand(c *proto.Client, m *proto.Message) bool {
	if len(config.CommandPrefix) == 0 || !strings.HasPrefix(m.Data, config.CommandPrefix) {
		return false
	}

	
	return true
}
