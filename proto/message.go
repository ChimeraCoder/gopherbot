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
	SenderName string
	SenderMask string
	Receiver   string
	Data       string
	Command    uint16
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
		m.Command = CPing
		m.Data = list[1][1:]
		return

	case "ERROR":
		m.Command = CError
		m.Data = list[1][1:]
		return
	}

	m.SenderMask = list[0][1:]
	idx := strings.Index(m.SenderMask, "!")

	if idx > -1 {
		m.SenderName = m.SenderMask[:idx]
		m.SenderMask = m.SenderMask[idx+1:]
	}

	m.Command = findType(list[1])
	m.Receiver = list[2]
	m.Data = strings.Join(list[3:], " ")

	if len(m.Data) > 0 && m.Data[0] == ':' {
		m.Data = m.Data[1:]
	}

	if len(m.Receiver) > 0 && m.Receiver[0] != '#' && m.Receiver[0] != '&' &&
		m.Receiver[0] != '!' && m.Receiver[0] != '+' {
		m.Receiver = m.SenderName
	}

	if m.Command == CPrivMsg {
		switch {
		case m.Data == "\x01VERSION\x01":
			m.Command = CCtcpVersion

		case strings.HasPrefix(m.Data, "\x01PING "):
			// \x01PING 1349825341 894301\x01
			m.Command = CCtcpPing
			m.Data = m.Data[6:]
		}
	}

	return
}

// findType attempts to parse a protocol ID from the input string.
// These come as 3-digit numbers or a string. For example: "001" or "NOTICE"
func findType(v string) uint16 {
	n, err := strconv.ParseUint(v, 10, 16)
	if err == nil {
		return uint16(n)
	}

	v = strings.ToUpper(v)

	switch v {
	case "ADMIN":
		return CAdmin
	case "AWAY":
		return CAway
	case "CONNECT":
		return CConnect
	case "DIE":
		return CDie
	case "ERROR":
		return CError
	case "INFO":
		return CInfo
	case "INVITE":
		return CInvite
	case "ISON":
		return CIsOn
	case "JOIN":
		return CJoin
	case "KICK":
		return CKick
	case "KILL":
		return CKill
	case "LINKS":
		return CLinks
	case "LIST":
		return CList
	case "LUSERS":
		return CLUsers
	case "MODE":
		return CMode
	case "MOTD":
		return CMOTD
	case "NAMES":
		return CNames
	case "NICK":
		return CNick
	case "NJOIN":
		return CNJoin
	case "NOTICE":
		return CNotice
	case "OPER":
		return COper
	case "PART":
		return CPart
	case "PASS":
		return CPass
	case "PING":
		return CPing
	case "PONG":
		return CPong
	case "PRIVMSG":
		return CPrivMsg
	case "QUIT":
		return CQuit
	case "REHASH":
		return CRehash
	case "RESTART":
		return CRestart
	case "SERVER":
		return CServer
	case "SERVICE":
		return CService
	case "SERVLIST":
		return CServList
	case "SQUERY":
		return CSQuery
	case "SQUIRT":
		return CSquirt
	case "SQUIT":
		return CSQuit
	case "STATS":
		return CStats
	case "SUMMON":
		return CSummon
	case "TIME":
		return CTime
	case "TOPIC":
		return CTopic
	case "TRACE":
		return CTrace
	case "USER":
		return CUser
	case "USERHOST":
		return CUserHost
	case "USERS":
		return CUsers
	case "VERSION":
		return CVersion
	case "WALLOPS":
		return CWAllOps
	case "WHO":
		return CWho
	case "WHOIS":
		return CWhoIs
	case "WHOWAS":
		return CWhoWas
	}

	return Unknown
}
