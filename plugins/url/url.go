// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package url

import (
	"bytes"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base

	// Each entry holds a regex pattern which should be excluded
	// from the url-title-lookup.
	exclude []*regexp.Regexp

	// Pattern which recognizes urls.
	url *regexp.Regexp
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "url")
	p.url = regexp.MustCompile(`\bhttps?\://[a-zA-Z0-9\-\.]+\.[a-zA-Z]+(\:[0-9]+)?(/\S*)?\b`)
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
		p.parseURL(c, m)
	})

	ini := p.LoadConfig()
	if ini != nil {
		s := ini.Section("exclude")
		list := s.List("url")
		p.exclude = make([]*regexp.Regexp, len(list))

		for i := range list {
			p.exclude[i], err = regexp.Compile(list[i])

			if err != nil {
				return
			}
		}
	}

	return
}

// parseURL looks for URL's embedded in incoming messages.
// If they are valid http[s] url's and not in the exclude list,
// we use them to fetch page titles from the internet.
func (p *Plugin) parseURL(c *proto.Client, m *proto.Message) {
	list := p.url.FindAllString(m.Data, -1)
	if len(list) == 0 {
		return
	}

	for _, url := range list {
		if !p.excluded(url) {
			go fetchTitle(c, m, url)
		}
	}
}

// excluded returns true if the given url is part of the exclusion list.
func (p *Plugin) excluded(url string) bool {
	for _, excl := range p.exclude {
		if excl.MatchString(url) {
			return true
		}
	}

	return false
}

// fetchTitle attempts to retrieve the title element for a given url.
func fetchTitle(c *proto.Client, m *proto.Message, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[url] Failed to fetch %s", url)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Printf("[url] Failed to read %s", url)
		return
	}

	body = bytes.ToLower(body)
	s := bytes.Index(body, []byte("<title>"))
	if s == -1 {
		return
	}

	body = body[s+7:]

	e := bytes.Index(body, []byte("</title>"))
	if e == -1 {
		e = len(body) - 1
	}

	body = bytes.TrimSpace(body[:e])

	c.PrivMsg(m.Receiver, "%s's link shows: %s",
		m.SenderName, html.UnescapeString(string(body)))
}
