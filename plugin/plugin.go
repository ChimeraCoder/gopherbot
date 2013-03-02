// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package plugin

import (
	"github.com/jteeuwen/ini"
	"github.com/jteeuwen/ircb/proto"
	"log"
	"path/filepath"
)

// PluginFunc represents a plugin constructor.
type PluginFunc func(string) Plugin

var (
	// List of registered plugin constructors.
	funcs []PluginFunc

	// List of registered plugins.
	plugins []Plugin
)

// Register registers a new plugin constructor.
// This is typically called in the init() function of a plugin package.
func Register(pf PluginFunc) { funcs = append(funcs, pf) }

// Load is called in the bot initialization and allows all registered
// plugins to initialize any necessary resources.
func Load(profile string, c *proto.Client) (err error) {
	log.Printf("Loading plugins...")

	for _, pf := range funcs {
		p := pf(profile)

		log.Printf("-> %s", p.Name())

		if err = p.Load(c); err != nil {
			return
		}

		plugins = append(plugins, p)
	}

	return
}

// Unload unloads all plugin resources.
func Unload(c *proto.Client) {
	log.Printf("Unloading plugins...")

	for _, p := range plugins {
		log.Printf("-> %s", p.Name())

		p.Unload(c)
	}
}

type Plugin interface {
	Load(*proto.Client) error
	Unload(*proto.Client)
	LoadConfig() *ini.File
	Name() string
	Profile() string
}

// Base represents a single plugin instance. It takes care of
// some basic housekeeping.
type Base struct {
	profile string
	name    string
	client  *proto.Client
}

// New creates a new plugin base with the given profile and name.
func New(profile, name string) *Base {
	p := new(Base)
	p.profile = profile
	p.name = name
	return p
}

func (p *Base) Name() string             { return p.name }
func (p *Base) Profile() string          { return p.profile }
func (p *Base) Load(*proto.Client) error { return nil }
func (p *Base) Unload(*proto.Client)     {}

// LoadConfig reads the ini configuration file for the given plugin.
// Returns nil if the file does not exist.
func (p *Base) LoadConfig() *ini.File {
	ini := ini.New()
	err := ini.Load(configPath(p.profile, p.name))

	if err != nil {
		return nil
	}

	return ini
}

// configPath returns the fully qualified path for the
// given plugin's configuration file.
func configPath(profile, name string) string {
	path := filepath.Join(profile, "plugins")
	path = filepath.Join(path, name)
	return filepath.Join(path, "config.ini")
}
