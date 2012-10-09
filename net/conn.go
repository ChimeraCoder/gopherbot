// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package net

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"time"
)

// Conn represents a single tcp connection.
type Conn struct {
	net.Conn
	reader *bufio.Reader
}

// Dial opens a conneciton to the given address.
// Optionally it runs in secure mode (TLS). But only if both key and cert
// contain valid paths to respective ssl-key and ssl-certificate files.
func Dial(address, key, cert string) (c *Conn, err error) {
	c = new(Conn)

	if len(key) > 0 && len(cert) > 0 {
		cfg := new(tls.Config)
		cfg.Rand = nil
		cfg.Time = time.Now
		cfg.ServerName = address

		c.Conn, err = tls.Dial("tcp", address, cfg)
	} else {
		c.Conn, err = net.Dial("tcp", address)
	}

	if err != nil {
		return
	}

	c.reader = bufio.NewReader(c.Conn)
	return
}

// Close closes the connection.
func (c *Conn) Close() (err error) {
	if c.Conn != nil {
		err = c.Conn.Close()
		c.Conn = nil
	}
	c.reader = nil
	return
}

// Write writes the given message to the underlying stream.
// It ensures the data does not exceed 512 bytes as this is the limit
// for IRC payloads. Any excess data is simply truncated.
func (c *Conn) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if c.Conn == nil {
		return 0, io.EOF
	}

	if len(p) >= 512 {
		p = p[:512]
	}

	if p[len(p)-1] != '\n' {
		p[len(p)-1] = '\n'
	}

	return c.Conn.Write(p)
}

func (c *Conn) Read(p []byte) (n int, err error) {
	if c.Conn == nil {
		return 0, io.EOF
	}

	return c.Conn.Read(p)
}

// ReadLine reads a single line from the underlying stream.
func (c *Conn) ReadLine() (data []byte, err error) {
	if c.Conn == nil {
		return nil, io.EOF
	}

	data, err = c.reader.ReadBytes('\n')
	if err != nil {
		return
	}

	data = bytes.TrimSpace(data[:len(data)-1])
	return
}
