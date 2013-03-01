// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package url

import (
	"bytes"
	"github.com/jteeuwen/ini"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var (
	// Each entry holds a regex pattern which should be excluded
	// from the url-title-lookup.
	exclude []*regexp.Regexp

	// Patern which recognizes urls.
	url = regexp.MustCompile(`\bhttps?\://[a-zA-Z0-9\-\.]+\.[a-zA-Z]+(\:[0-9]+)?(/\S*)?\b`)
)

// Init initializes the plugin. it loads configuration data and binds
// commands and protocol handlers.
func Init(profile string, c *proto.Client) {
	log.Println("Initializing: url")

	ini := ini.New()
	err := ini.Load(plugin.ConfigPath(profile, "url"))

	if err != nil {
		log.Fatal(err)
	}

	s := ini.Section("exclude")
	list := s.List("url")
	exclude = make([]*regexp.Regexp, len(list))

	for i := range list {
		exclude[i], err = regexp.Compile(list[i])
		if err != nil {
			log.Fatalf("- Invalid pattern: %s", list[i])
		}
	}

	c.Bind(proto.CmdPrivMsg, parseURL)
}

// parseURL looks for URL's embedded in incoming messages.
// If they are valid http[s] url's and not in the exclude list,
// we use them to fetch page titles from the internet.
func parseURL(c *proto.Client, m *proto.Message) {
	list := url.FindAllString(m.Data, -1)
	if len(list) == 0 {
		return
	}

	for _, url := range list {
		if !excluded(url) {
			go fetchTitle(c, m, url)
		}
	}
}

// excluded returns true if the given url is part of the exclusion list.
func excluded(url string) bool {
	for _, excl := range exclude {
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
