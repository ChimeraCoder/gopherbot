// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package irc

// User represents a single irc user.
type User struct {
	Nickname string
	Hostmask string
	LastMsg  int64
}
