// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ipintel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/chimeracoder/gopherbot/cmd"
	"github.com/chimeracoder/gopherbot/plugin"
	"github.com/chimeracoder/gopherbot/proto"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var regMibbit = regexp.MustCompile(`^[a-fA-F0-9]{8}$`)

const url = "http://api.neustar.biz/ipi/std/v1/ipinfo/%s?apikey=%s&sig=%s&format=json"

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "ipintel")
	return p
}

func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	ini := p.LoadConfig()
	if ini == nil {
		log.Fatalf("[ipintel] No configuration found.")
		return
	}

	key := ini.Section("api").S("key", "")
	if len(key) == 0 {
		log.Fatalf("[ipintel] No API key found.")
		return
	}

	shared := ini.Section("api").S("shared", "")
	if len(shared) == 0 {
		log.Fatalf("[ipintel] No API shared secret found.")
		return
	}

	drift := ini.Section("api").I64("drift", 0)
	if len(shared) == 0 {
		log.Fatalf("[ipintel] No API shared secret found.")
		return
	}

	w := new(cmd.Command)
	w.Name = "loc"
	w.Description = "Fetch geo-location data for the given IP address."
	w.Restricted = false
	w.Params = []cmd.Param{
		{Name: "ip", Description: "IPv4 address to look up", Pattern: cmd.RegIPv4},
	}
	w.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		hash := md5.New()
		stamp := fmt.Sprintf("%d", time.Now().UTC().Unix()+drift)
		io.WriteString(hash, key+shared+stamp)

		sig := fmt.Sprintf("%x", hash.Sum(nil))
		target := fmt.Sprintf(url, cmd.Params[0].Value, key, sig)

		resp, err := http.Get(target)
		if err != nil {
			log.Printf("[ipintel]: %v", err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Printf("[ipintel]: %v", err)
			return
		}

		var data Response
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Printf("[ipintel]: %v", err)
			return
		}

		inf := data.IPInfo
		mapsURL := fmt.Sprintf("https://maps.google.com/maps?q=%f,%f",
			inf.Location.Latitude, inf.Location.Longitude)

		c.PrivMsg(m.Receiver,
			"%s: %s (%s), Network org.: %s, Carrier: %s, TLD: %s, SLD: %s. "+
				"Location: %s/%s/%s/%s (%f, %f). Postalcode: %s, Timezone: %d, %s",
			m.SenderName,

			inf.IPAddress, inf.IPType,
			inf.Network.Organization,
			inf.Network.Carrier,
			inf.Network.Domain.TLD,
			inf.Network.Domain.SLD,

			inf.Location.Continent,
			inf.Location.CountryData.Name,
			inf.Location.StateData.Name,
			inf.Location.CityData.Name,
			inf.Location.Latitude,
			inf.Location.Longitude,
			inf.Location.CityData.PostalCode,
			inf.Location.CityData.TimeZone,

			mapsURL,
		)
	}

	cmd.Register(w)

	w = new(cmd.Command)
	w.Name = "mibbit"
	w.Description = "Resolve a mibbit address to a real IP address."
	w.Restricted = false
	w.Params = []cmd.Param{
		{Name: "hex", Description: "Mibbit hex string", Pattern: regMibbit},
	}
	w.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		hex := cmd.Params[0].Value

		var ip [4]uint64
		var err error

		ip[0], err = strconv.ParseUint(hex[:2], 16, 8)
		if err != nil {
			c.PrivMsg(m.Receiver, "%s: invalid mibbit address.", m.SenderName)
			return
		}

		ip[1], err = strconv.ParseUint(hex[2:4], 16, 8)
		if err != nil {
			c.PrivMsg(m.Receiver, "%s: invalid mibbit address.", m.SenderName)
			return
		}

		ip[2], err = strconv.ParseUint(hex[4:6], 16, 8)
		if err != nil {
			c.PrivMsg(m.Receiver, "%s: invalid mibbit address.", m.SenderName)
			return
		}

		ip[3], err = strconv.ParseUint(hex[6:], 16, 8)
		if err != nil {
			c.PrivMsg(m.Receiver, "%s: invalid mibbit address.", m.SenderName)
			return
		}

		address := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
		names, err := net.LookupAddr(address)

		if err != nil || len(names) == 0 {
			c.PrivMsg(m.Receiver, "%s: %s is %s", m.SenderName, hex, address)
		} else {
			c.PrivMsg(m.Receiver, "%s: %s is %s / %s",
				m.SenderName, hex, address, names[0])
		}
	}

	cmd.Register(w)

	return
}
