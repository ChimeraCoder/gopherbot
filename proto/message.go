// This file is subject to a 1-clause BSD license.
// Its Datas can be found in the enclosed LICENSE file.

package proto

import (
	"io"
	"strconv"
	"strings"
)

// Message is a parsed incoming message.
type Message struct {
	SenderName string // Nickname of sender.
	SenderMask string // Hostmask of sender.
	Receiver   string // Target of message. Can be a user (our bot) or channel.
	Data       string // Message payload.
	Command    uint16 // Command identifier: type of message.
}

// FromChannel returns true if this message came from a channel context
// instead of a user or service.
func (m *Message) FromChannel() bool {
	if len(m.Receiver) == 0 {
		return false
	}

	c := m.Receiver[0]
	return c == '#' || c == '&' || c == '!' || c == '+'
}

// parseMessage parses a message from the given data.
func parseMessage(data string) (m *Message, err error) {
	if len(data) == 0 {
		return nil, io.EOF
	}

	m = new(Message)
	m.Command = Unknown
	m.Data = data

	list := strings.Split(data, " ")

	switch list[0] {
	case "PING":
		m.Command = CmdPing
		m.Data = list[1][1:]
		return

	case "ERROR":
		m.Command = CmdError
		m.Data = list[1][1:]
		return
	}

	if len(list) < 3 {
		return
	}

	m.SenderMask = list[0][1:]
	m.Command = findType(list[1])
	m.Receiver = list[2]
	m.Data = strings.Join(list[3:], " ")

	idx := strings.Index(m.SenderMask, "!")
	if idx > -1 {
		m.SenderName = m.SenderMask[:idx]
		m.SenderMask = m.SenderMask[idx+1:]
	}

	if len(m.Data) > 0 && m.Data[0] == ':' {
		m.Data = m.Data[1:]
	}

	return
}

// findType attempts to parse a command or reply type from the input string.
// These come as 3-digit numbers or a string. For example: "001" or "NOTICE"
func findType(v string) uint16 {
	n, err := strconv.ParseUint(v, 10, 16)
	if err == nil {
		return uint16(n)
	}

	v = strings.ToUpper(v)

	switch v {
	case "ADMIN":
		return CmdAdmin
	case "AWAY":
		return CmdAway
	case "CONNECT":
		return CmdConnect
	case "DIE":
		return CmdDie
	case "ERROR":
		return CmdError
	case "INFO":
		return CmdInfo
	case "INVITE":
		return CmdInvite
	case "ISON":
		return CmdIsOn
	case "JOIN":
		return CmdJoin
	case "KICK":
		return CmdKick
	case "KILL":
		return CmdKill
	case "LINKS":
		return CmdLinks
	case "LIST":
		return CmdList
	case "LUSERS":
		return CmdLUsers
	case "MODE":
		return CmdMode
	case "MOTD":
		return CmdMOTD
	case "NAMES":
		return CmdNames
	case "NICK":
		return CmdNick
	case "NJOIN":
		return CmdNJoin
	case "NOTICE":
		return CmdNotice
	case "OPER":
		return CmdOper
	case "PART":
		return CmdPart
	case "PASS":
		return CmdPass
	case "PING":
		return CmdPing
	case "PONG":
		return CmdPong
	case "PRIVMSG":
		return CmdPrivMsg
	case "QUIT":
		return CmdQuit
	case "REHASH":
		return CmdRehash
	case "RESTART":
		return CmdRestart
	case "SERVER":
		return CmdServer
	case "SERVICE":
		return CmdService
	case "SERVLIST":
		return CmdServList
	case "SQUERY":
		return CmdSQuery
	case "SQUIRT":
		return CmdSquirt
	case "SQUIT":
		return CmdSQuit
	case "STATS":
		return CmdStats
	case "SUMMON":
		return CmdSummon
	case "TIME":
		return CmdTime
	case "TOPIC":
		return CmdTopic
	case "TRACE":
		return CmdTrace
	case "USER":
		return CmdUser
	case "USERHOST":
		return CmdUserHost
	case "USERS":
		return CmdUsers
	case "VERSION":
		return CmdVersion
	case "WALLOPS":
		return CmdWAllOps
	case "WHO":
		return CmdWho
	case "WHOIS":
		return CmdWhoIs
	case "WHOWAS":
		return CmdWhoWas
	}

	return Unknown
}
