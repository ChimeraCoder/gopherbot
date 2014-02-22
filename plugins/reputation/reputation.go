// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package reputation

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"log"
	"os"
	"regexp"
	"strconv"
)

type RepChange struct {
	entity string
	delta  int
}

var reputation map[string]int

var reputationChanges chan RepChange

var red redis.Conn

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
	p.sexpr = regexp.MustCompile(`\((\+\+|--) (.*?)\)`)
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
			go scoreReputation(c, m, sexpr)
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
func scoreReputation(c *proto.Client, m *proto.Message, match []string) {
	entity := match[2]
	action := match[1]
	action_words := " is in limbo"
	switch action {
	case "++":
		log.Printf("incrementing %s", entity)
		_, err := red.Do("INCR", entity)
		if err != nil {
			log.Print(err)
		}
		action_words = "gained 1 rep"
	case "--":
		log.Printf("decrementing %s", entity)
		_, err := red.Do("DECR", entity)
		if err != nil {
			log.Print(err)
		}
		action_words = "lost 1 rep"
	case "rep":
		action_words = "has rep"

	default:
		log.Printf("action %s not supported", action)
		return
	}

	rep, err := red.Do("GET", entity)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	rep_b, ok := rep.([]byte)
	if !ok {
		log.Printf("ERROR: not a byte slice type: %v", rep)
		return
	}
	log.Printf("Fetched %s", string(rep_b))
	rep_i, err := strconv.Atoi(string(rep_b))
	if err != nil {
		log.Printf("Error converting to integer %s", err)
		return
	}

	c.PrivMsg(m.Receiver, "%s %s! rep: %d", entity, action_words, rep_i)
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
