// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package reputation

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"log"
	"os"
	"regexp"
	"strings"
)

type RepChange struct {
	entity string
	delta  int
}

var reputation map[string]int

var reputationChanges chan RepChange

var red redis.Conn

const WHOIS_DB = "2"
const WHOIS_SUFFIX = "-whois" // TODO remove this hack

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base

	// Each entry holds a regex pattern which should be excluded
	// from the url-title-lookup.
	exclude []*regexp.Regexp

	// Pattern which recognizes s-expressions.
	// TODO use proper parser
	sexpr *regexp.Regexp
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "url")
	p.sexpr = regexp.MustCompile(`(\((whois) (.*?)\)|\((is) (.*?) (.*?)\))`)
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
		p.parseSexpr(c, m)
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

// parseSexpr looks for s-expressions embedded in incoming messages.
// If they are valid s-expressions and not in the exclude list,
// we use them to fetch page titles from the internet.
func (p *Plugin) parseSexpr(c *proto.Client, m *proto.Message) {
	list := p.sexpr.FindAllStringSubmatch(m.Data, -1)
	if len(list) == 0 {
		return
	}

	for _, sexpr := range list {
		if !p.excluded(sexpr[0]) {
			go whoIs(c, m, sexpr)
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
func whoIs(c *proto.Client, m *proto.Message, match []string) {
	log.Printf("Whois: %+v", match)
	entity := strings.ToLower(match[3])
	action := match[2]
	response := "no idea who or what that is"
	switch action {
	case "whois":
		descriptor_b, err := red.Do("GET", entity+WHOIS_SUFFIX)
		if err != nil {
			log.Print(err)
			break
		}

		// Will be nil if not yet in the database
		if descriptor_b == nil {
			break
		}
		descriptor, ok := descriptor_b.([]byte)
		if !ok {
			fmt.Printf("ERROR: not a byte slice type: %+v", descriptor_b)
			break
		}

		response = fmt.Sprintf("%s is %s", entity, string(descriptor)+"...")
	default:
		if match[4] != "is" {
			log.Printf("action %s not supported", match[3])
			return
		}
		entity = match[5]
		descriptor := match[6]
		_, err := red.Do("APPEND", entity+WHOIS_SUFFIX, descriptor+", ")
		if err != nil {
			log.Print(err)
			return
		}
		response = fmt.Sprintf("%s is %s", entity, descriptor)
	}

	c.PrivMsg(m.Receiver, response)
}

func init() {
	var err error
	red, err = redis.Dial(os.Getenv("REDIS_NETWORK"), os.Getenv("REDIS_ADDRESS"))
	if err != nil {
		log.Printf("ERROR: Failed to connect to redis database - reputation plugin will NOT load")
		panic(err)
	}
	_, err = red.Do("AUTH", os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		panic(err)
	}
}
