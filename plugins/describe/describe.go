// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package describe

import (
	"fmt"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"log"
	"regexp"
)

//This regex will check if a URL points to a Twitter status
var descriptionRegex = regexp.MustCompile(`ACTION is (.*)`)

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base

	// Each entry holds a regex pattern which should be excluded
	exclude []*regexp.Regexp

	// Pattern which recognizes descriptions.
	description *regexp.Regexp
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "description")
	p.description = descriptionRegex
	return p
}

// Init initializes the plugin. it loads configuration data and binds
// commands and protocol handlers.
func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	c.Bind(proto.CmdPrivMsg, func(c *proto.Client, m *proto.Message) {
		p.parseDescription(c, m)
	})

	ini := p.LoadConfig()
	if ini == nil {
		return
	}

	s := ini.Section("exclude")
	list := s.List("description")
	p.exclude = make([]*regexp.Regexp, len(list))

	for i := range list {
		p.exclude[i], err = regexp.Compile(list[i])

		if err != nil {
			return
		}
	}

	return
}

// parseURL looks for descriptions in incoming messages.
func (p *Plugin) parseDescription(c *proto.Client, m *proto.Message) {
	log.Printf("Checking %s", m.Data)
	list := p.description.FindStringSubmatch(m.Data)
	if len(list) == 0 {
		return
	}

	description := list[len(list)-1]
	log.Printf("Found description: %s", description)
	c.PrivMsg(m.Receiver, fmt.Sprintf("%s is %s", m.SenderName, description))
}
