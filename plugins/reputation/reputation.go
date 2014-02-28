// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package reputation

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	/* \S means NOT whitespace, ?: is a non-capture group */
	p.sexpr = regexp.MustCompile(`\((\+\+|--|rep|1\+|1-|1\?)[\s]+([\S]+?)(?:[\s]+?)\)`)
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

/* @todo check "reserved" keywords such as top/bottom */
func incrementReputation(c *proto.Client, m *proto.Message, entity string) {
	log.Printf("incrementing %s", entity)
	rep, err := red.Do("ZINCRBY", "reputation", "1", entity)
	if err != nil {
		log.Print(err)
		return
	}
	c.PrivMsg(m.Receiver, "%s gained 1 rep! rep: %s",
		entity, string(rep.([]byte)))
}

func decrementReputation(c *proto.Client, m *proto.Message, entity string) {
	log.Printf("decrementing %s", entity)
	rep, err := red.Do("ZINCRBY", "reputation", "-1", entity)
	if err != nil {
		log.Print(err)
		return
	}
	c.PrivMsg(m.Receiver, "%s lost 1 rep! rep: %s",
		entity, string(rep.([]byte)))
}

func checkReputation(c *proto.Client, m *proto.Message, entity string) {
	log.Printf("checking %s", entity)
	rep, err := red.Do("ZSCORE", "reputation", entity)
	if err != nil {
		log.Print(err)
		return
	}
	if rep == nil {
		c.PrivMsg(m.Receiver, "never heard of %s", entity)
		return
	}

	c.PrivMsg(m.Receiver, "%s has rep: %s", entity, string(rep.([]byte)))
}

func printResponseReputationList(c *proto.Client, m *proto.Message,
	resp interface{}) {
	values, err := redis.Values(resp, nil)
	if err != nil {
		log.Print(err)
		return
	}
	var reps []struct {
		Name  string
		Score int
	}
	err = redis.ScanSlice(values, &reps)
	for i, rep := range reps {
		c.PrivMsg(m.Receiver, "(%d) %-10s: %3d", i, rep.Name, rep.Score)
	}
}

func listTopReputation(c *proto.Client, m *proto.Message) {
	resp, err := red.Do("ZREVRANGE", "reputation", "0", "4", "WITHSCORES")
	if err != nil {
		log.Print(err)
		return
	}
	c.PrivMsg(m.Receiver, "Top Reputations:")
	printResponseReputationList(c, m, resp)
}

func listBotReputation(c *proto.Client, m *proto.Message) {
	resp, err := red.Do("ZRANGE", "reputation", "0", "4", "WITHSCORES")
	if err != nil {
		log.Print(err)
		return
	}
	c.PrivMsg(m.Receiver, "Bottom Reputations:")
	printResponseReputationList(c, m, resp)
}

func scoreReputation(c *proto.Client, m *proto.Message, match []string) {
	entity := strings.ToLower(match[2])
	action := match[1]

	switch action {
	case "1+":
		fallthrough
	case "++":
		incrementReputation(c, m, entity)

	case "1-":
		fallthrough
	case "--":
		decrementReputation(c, m, entity)

	case "1?":
		log.Printf("random %s", entity)
		if rand.Intn(2) == 0 {
			incrementReputation(c, m, entity)
		} else {
			decrementReputation(c, m, entity)
		}

	case "rep":
		strconv.Atoi("lol")
		if entity == "top" {
			listTopReputation(c, m)
		} else if entity == "bot" {
			listBotReputation(c, m)
		} else {
			checkReputation(c, m, entity)
		}
	default:
		log.Printf("action %s not supported", action)
		return
	}

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
	rand.Seed(time.Now().UTC().UnixNano())
}
