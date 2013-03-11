// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ipintel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/jteeuwen/ircb/cmd"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

		c.PrivMsg(m.Receiver,
			"%s: %s (%s), Network org.: %s, Carrier: %s, "+
				"Conn. type: %s. Conn. speed: %s, Routing: %s, TLD: %s, SLD: %s. "+
				"Location: %s/%s/%s/%s (%f, %f). Postalcode: %s, Timezone: %d",
			m.SenderName,

			inf.IPAddress, inf.IPType,
			inf.Network.Organization,
			inf.Network.Carrier,
			inf.Network.ConnectionType,
			inf.Network.LineSpeed,
			inf.Network.RoutingType,
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
		)
	}

	cmd.Register(w)

	return
}
