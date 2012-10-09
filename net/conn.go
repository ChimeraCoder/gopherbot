// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package net

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net"
)

// Conn represents a single tcp connection.
type Conn struct {
	net.Conn
	reader *bufio.Reader
}

// Dial opens a connection to the given address.
// Optionally it runs in secure mode (TLS). But only if both key and cert
// contain valid paths to respective x509 key and certificate files.
func Dial(address, cert, key string) (c *Conn, err error) {
	c = new(Conn)

	log.Printf("Connecting to %s...", address)

	if len(key) > 0 && len(cert) > 0 {
		var cfg tls.Config
		//cfg.InsecureSkipVerify = true
		cfg.Certificates = make([]tls.Certificate, 1)
		cfg.Certificates[0], err = tls.LoadX509KeyPair(cert, key)

		if err != nil {
			return
		}

		c.Conn, err = tls.Dial("tcp", address, &cfg)
		if err != nil {
			return
		}

		log.Println("Secure connection established.")

		state := c.Conn.(*tls.Conn).ConnectionState()
		log.Println("TLS Handshake: ", state.HandshakeComplete)
		log.Println("TLS Mutual: ", state.NegotiatedProtocolIsMutual)
		log.Println("TLS Certificates:")

		for _, v := range state.PeerCertificates {
			log.Printf(" - %v", v.Subject)
		}
	} else {
		if c.Conn, err = net.Dial("tcp", address); err != nil {
			return
		}

		log.Println("Connection established.")
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
	if len(p) == 0 || c.Conn == nil {
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
