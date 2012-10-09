// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/blah/net"
	"github.com/jteeuwen/blah/proto"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cfg, conn, client := setup()
	defer shutdown(cfg, conn, client)

	// Perform handshake.
	log.Printf("Performing handshake...")
	client.Login(cfg.Nickname)
	client.Nick(cfg.Nickname, cfg.NickservPassword)

	// Main data loop.
	log.Printf("Entering data loop...")
	for {
		line, err := conn.ReadLine()

		if err != nil {
			break
		}

		client.Read(string(line))
	}
}

// setup initializes the application.
func setup() (*Config, *net.Conn, *proto.Client) {
	cfg := parseArgs()

	// Open connection to server.
	log.Printf("Connecting to %s...", cfg.Address)
	conn, err := net.Dial(cfg.Address, cfg.SSLKey, cfg.SSLCert)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Dial: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Connection established.")

	// Create protocol handler.
	client := proto.NewClient(func(p []byte) error {
		log.Printf("< %s", p)
		_, err := conn.Write(p)
		return err
	})

	Bind(client)
	return cfg, conn, client
}

// shutdown cleans up our mess.
func shutdown(cfg *Config, conn *net.Conn, client *proto.Client) {
	log.Printf("Shutting down.")
	client.Quit(cfg.QuitMessage)
	client.Close()
	conn.Close()
}

// parseArgs reads and verfies commandline arguments.
// It loads and returns a configuration object.
func parseArgs() *Config {
	config := flag.String("c", "", "Path to bot configuraiton file.")
	version := flag.Bool("v", false, "Display version information.")

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if len(*config) == 0 {
		fmt.Fprintf(os.Stderr, "Missing configuraiton file.\n")
		flag.Usage()
		os.Exit(1)
	}

	var c Config
	err := c.Load(filepath.Clean(*config))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Load config: %v\n", err)
		os.Exit(1)
	}

	return &c
}
