// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package url

import (
	"bytes"
	"github.com/ChimeraCoder/anaconda"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//This regex will check if a URL points to a Twitter status
var twitterUrlRegex = regexp.MustCompile(`https?:\/\/(www\.)?twitter.com\/[A-Za-z0-9]*\/status\/([0-9]+)`)

var api anaconda.TwitterApi

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
	if ini == nil {
		return
	}

	s := ini.Section("exclude")
	list := s.List("url")
	p.exclude = make([]*regexp.Regexp, len(list))

	for i := range list {
		p.exclude[i], err = regexp.Compile(list[i])

		if err != nil {
			return
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
		//TODO make this less hackny
		if twitterUrlRegex.MatchString(url) {
			go fetchTweet(c, m, url)

		} else if !p.excluded(url) {
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
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
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

// fetchTweet attempts to retrieve the tweet associated with a given url.
func fetchTweet(c *proto.Client, m *proto.Message, url string) {
	id, err := strconv.ParseInt(twitterUrlRegex.FindStringSubmatch(url)[2], 10, 64)
	if err != nil {
		c.PrivMsg(m.Receiver, "error parsing tweet :(")
		log.Print("error parsing tweet for %s: %v", url, err)
		fetchTitle(c, m, url)
		return
	}
	tweet, err := api.GetTweet(id, nil)
	if err != nil {
		log.Print("error parsing tweet for %s: %v", url, err)
		fetchTitle(c, m, url)
		return
	}
	c.PrivMsg(m.Receiver, "%s's tweet shows: %s",
		m.SenderName, html.UnescapeString(tweet.Text))
}

func init() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api = anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}
