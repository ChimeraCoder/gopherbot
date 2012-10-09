// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package proto

import (
	"fmt"
	"github.com/jteeuwen/blah/irc"
)

// ReadHandler represents a client protocol event handler.
// It accepts the client instance and the message object that was created.
type ReadHandler func(*Client, *Message)

// WriteHandler will be called whenever a new message is being created
// by the client type. An implementation of this signature can forward
// the data to an underlying network connection.
type WriteHandler func(data []byte) error

// Client wraps an io.Writer and exposes IRC client protocol methods.
type Client struct {
	writer WriteHandler              // Write handler.
	events map[ProtoId][]ReadHandler // Bound protocol event handlers.
}

// NewClient creates a new client for the given writer.
func NewClient(writer WriteHandler) *Client {
	c := new(Client)
	c.writer = writer
	c.events = make(map[ProtoId][]ReadHandler)
	return c
}

// Close cleans up the client.
func (c *Client) Close() (err error) {
	c.writer = nil
	c.events = nil
	return
}

// Read attempts to parse the given data into a Message object.
// It fires the approriate protocol handlers when applicable.
func (c *Client) Read(data string) (err error) {
	msg, err := parseMessage(data)
	if err != nil {
		return
	}

	m, ok := c.events[PIDUnknown]
	if ok {
		for _, f := range m {
			f(c, msg)
		}
	}

	if msg.Type == PIDUnknown {
		return // Already called these handlers.
	}

	if m, ok = c.events[msg.Type]; !ok {
		return
	}

	for _, f := range m {
		f(c, msg)
	}

	return
}

// Bind binds the given read handler to the specified protocol type.
// The handler is called whenever a message of the given type is received.
// There can be multiple handlers for a single protocol message type.
//
// Binding to PIDUnknown, will trigger the given handler on _every_
// incoming message. This can be useful if you just wish to agregate all
// incoming data, regardless of its type. 
func (c *Client) Bind(proto ProtoId, ch ReadHandler) {
	c.events[proto] = append(c.events[proto], ch)
}

// Raw sends the given message data to the specified writer.
func (c *Client) Raw(f string, argv ...interface{}) error {
	return c.writer([]byte(fmt.Sprintf("%s\n", fmt.Sprintf(f, argv...))))
}

// Login performs the initial connection handshake.
// It should usually be followed directly with a call to Client.Nick().
func (c *Client) Login(username string) error {
	return c.Raw("USER %s server %s :%s user", username, username, username)
}

// Privmsg sends the specified message to the given target.
func (c *Client) Privmsg(target, f string, argv ...interface{}) error {
	return c.Raw("PRIVMSG %s :%s", target, fmt.Sprintf(f, argv...))
}

// Notice sends the specifid notice to the given target.
func (c *Client) Notice(target, f string, argv ...interface{}) error {
	return c.Raw("NOTICE %s :%s", target, fmt.Sprintf(f, argv...))
}

// Quit quits from the server, optionally with the given quit message.
func (c *Client) Quit(f string, argv ...interface{}) error {
	f = fmt.Sprintf(f, argv...)

	if len(f) > 0 {
		return c.Raw("QUIT %s", f)
	}

	return c.Raw("QUIT")
}

// Pong sends the given payload as a response to the PING message.
func (c *Client) Pong(payload string) error {
	return c.Raw("PONG %s", payload)
}

// Nick changes the current nickname and optionally identifies with
// the given password.
func (c *Client) Nick(name, pass string) error {
	err := c.Raw("NICK " + name)
	if err != nil {
		return err
	}

	if len(pass) > 0 {
		err = c.Privmsg("nickserv", "IDENTIFY "+pass)
	}
	return err
}

// Join joins the given channels.
func (c *Client) Join(channels []*irc.Channel) (err error) {
	for _, ch := range channels {
		if err = c.Raw("CS INVITE %s", ch.Name); err != nil {
			return
		}

		if len(ch.Key) > 0 {
			err = c.Raw("JOIN %s %s", ch.Name, ch.Key)
		} else {
			err = c.Raw("JOIN %s", ch.Name)
		}

		if err != nil {
			return
		}

		if len(ch.ChanservPassword) > 0 {
			// FIXME(jimt): Ensure this is correct.
			// Do we need to send the channel name?
			err = c.Privmsg("chanserv", "IDENTIFY "+ch.ChanservPassword)
			if err != nil {
				return
			}
		}
	}

	return
}

// Part leaves the given channels.
func (c *Client) Part(channels []*irc.Channel) (err error) {
	for _, ch := range channels {
		if err = c.Raw("PART %s", ch.Name); err != nil {
			return
		}
	}
	return
}

// Mode changes the mode for the given target.
// Optionally with the given argument.
func (c *Client) Mode(target, mode, arg string) error {
	if len(arg) > 0 {
		return c.Raw("MODE %s %s %s", target, mode, arg)
	}
	return c.Raw("MODE %s %s", target, mode)
}

// Topic sets the given topic. This assumes we are in a channel context.
func (c *Client) Topic(target, topic string) error {
	return c.Raw("TOPIC %s :%s", target, topic)
}

// Invite invites the given nick to the specified target channel.
// Optionally with the given message.
func (c *Client) Invite(nick, target, message string) error {
	return c.Raw(":%s INVITE %s :%s", nick, target, message)
}

// Kick kicks the given target from the specified channel.
// Optionally with the given reason.
func (c *Client) Kick(channel, target, reason string) error {
	if len(reason) > 0 {
		return c.Raw("KICK %s %s :%s", channel, target, reason)
	}
	return c.Raw("KICK %s %s", channel, target)
}
