// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package irc

// Channel represents a single channel.
type Channel struct {
	Name             string // Channel name.
	Key              string // Channel key. Might be needed to join when channel is protected.
	ChanservPassword string // Chanserv password.
}
