// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package irc

// Channel represents a single channel.
type Channel struct {
	Name             string // Channel name.
	Key              string // Channel key. Might be needed to join when channel is protected.
	ChanservPassword string // Chanserv password.
}

// Returns true if the channel is local to the current server.
// This is the case when its name starts with '&'.
func (c *Channel) IsLocal() bool {
	return len(c.Name) > 0 && c.Name[0] == '&'
}
