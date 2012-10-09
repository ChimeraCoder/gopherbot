// This file is subject to a 1-clause BSD license.
// Its Datas can be found in the enclosed LICENSE file.

package proto

import (
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	regError  = regexp.MustCompile(`^ERROR :(.+)$`)
	regPing   = regexp.MustCompile(`^PING :(.+)$`)
	regNotice = regexp.MustCompile(`^NOTICE .+ :(.+)$`)
	regSrv1   = regexp.MustCompile(`^:(.+) (\d+) (.+) (.+) :(.+)$`)
	regSrv2   = regexp.MustCompile(`^:(.+) (\d+) (.+) :(.+)$`)
	regMsg    = regexp.MustCompile(`^:(.+)!(.+) (.+) (.+) :(.*)$`)
)

// Message is a parsed incoming message.
type Message struct {
	Server     string
	SenderName string
	SenderMask string
	Receiver   string
	Data       string
	Type       ProtoId
}

// parseMessage parses a message from the given data.
func parseMessage(data string) (m *Message, err error) {
	if len(data) == 0 {
		return nil, io.EOF
	}

	m = new(Message)
	m.Type = PIDUnknown
	m.Data = data

	match := regMsg.FindStringSubmatch(data)
	if len(match) > 0 {
		m.SenderName = match[1]
		m.SenderMask = match[2]
		m.Type = findType(match[3])
		m.Receiver = match[4]

		if m.Receiver[0] != '#' && m.Receiver[0] != '&' &&
			m.Receiver[0] != '!' && m.Receiver[0] != '+' {
			m.Receiver = m.SenderName
		}

		m.Data = strings.TrimSpace(match[5])

		if m.Data == "\x01VERSION\x01" {
			m.Type = PIDVersion
		}

		return
	}

	match = regError.FindStringSubmatch(data)
	if len(match) > 0 {
		m.Type = PIDError
		m.Data = strings.TrimSpace(match[1])
		return
	}

	match = regPing.FindStringSubmatch(data)
	if len(match) > 0 {
		m.Type = PIDPing
		m.Data = strings.TrimSpace(match[1])
		return
	}

	match = regNotice.FindStringSubmatch(data)
	if len(match) > 0 {
		m.Type = PIDNotice
		m.Data = strings.TrimSpace(match[1])
		return
	}

	match = regSrv1.FindStringSubmatch(data)
	if len(match) > 0 {
		m.Server = match[1]
		m.Type = findType(match[2])
		m.SenderName = match[3]
		m.Receiver = match[4]

		if m.SenderName == "*" {
			m.SenderName = m.Server
		}

		if len(match) > 5 && len(match[5]) > 0 {
			m.Data = strings.TrimSpace(match[5])
		}
		return
	}

	match = regSrv2.FindStringSubmatch(data)
	if len(match) > 0 {
		m.Server = match[1]
		m.Type = findType(match[2])
		m.SenderName = m.Server
		m.Receiver = match[3]

		if len(match) > 4 && len(match[4]) > 0 {
			m.Data = strings.TrimSpace(match[4])
		}
	}

	return
}

// findType attempts to parse a protocol ID from the input string.
// These come as 3-digit numbers or a string. For example: "001" or "NOTICE"
func findType(v string) ProtoId {
	n, err := strconv.ParseUint(v, 10, 16)
	if err == nil {
		return ProtoId(n)
	}

	v = strings.ToUpper(v)

	switch v {
	case "NOTICE":
		return PIDNotice
	case "PRIVMSG":
		return PIDPrivMsg
	case "QUIT":
		return PIDQuit
	case "JOIN":
		return PIDJoin
	case "PART":
		return PIDPart
	case "KICK":
		return PIDKick
	case "NICK":
		return PIDNick
	case "PING":
		return PIDPing
	case "ERROR":
		return PIDError
	case "MODE":
		return PIDMode
	}

	return PIDUnknown
}
