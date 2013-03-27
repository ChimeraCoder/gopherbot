// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package weather

import (
	"encoding/json"
	"fmt"
	"github.com/jteeuwen/ircb/cmd"
	"github.com/jteeuwen/ircb/plugin"
	"github.com/jteeuwen/ircb/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const _url = `http://free.worldweatheronline.com/feed/weather.ashx?format=json&num_of_days=2&q=%s&key=%s`

func init() { plugin.Register(New) }

type Plugin struct {
	*plugin.Base
}

func New(profile string) plugin.Plugin {
	p := new(Plugin)
	p.Base = plugin.New(profile, "weather")
	return p
}

func (p *Plugin) Load(c *proto.Client) (err error) {
	err = p.Base.Load(c)
	if err != nil {
		return
	}

	ini := p.LoadConfig()
	if ini == nil {
		log.Fatalf("[weather] No API key found.")
		return
	}

	key := ini.Section("api").S("key", "")
	if len(key) == 0 {
		log.Fatalf("[weather] No API key found.")
		return
	}

	w := new(cmd.Command)
	w.Name = "weather"
	w.Description = "Fetch the current weather for a given location"
	w.Restricted = false
	w.Params = []cmd.Param{
		{Name: "location", Description: "Name of the city/town for the forecast", Pattern: cmd.RegAny},
	}

	w.Execute = func(cmd *cmd.Command, c *proto.Client, m *proto.Message) {
		location := url.QueryEscape(cmd.Params[0].Value)
		resp, err := http.Get(fmt.Sprintf(_url, location, key))
		if err != nil {
			return
		}

		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			return
		}

		var wd WeatherData
		err = json.Unmarshal(data, &wd)
		if err != nil {
			return
		}

		if len(wd.Data.Request) == 0 || len(wd.Data.Conditions) == 0 {
			c.PrivMsg(m.Receiver, "%s: No weather data for %q",
				m.SenderName, cmd.Params[0].Value)
			return
		}

		wr := wd.Data.Request[0]
		wc := wd.Data.Conditions[0]

		c.PrivMsg(m.Receiver,
			"%s, weather in %s: %s°C/%s°F/%.2f°K, %s, cloud cover: %s%%, humidity: %s%%, wind: %skph/%smph from %s, pressure: %s mb, visibility: %s km",
			m.SenderName, wr.Query,
			wc.TempC, wc.TempF, wd.TempK(), codeName(wc.WeatherCode),
			wc.CloudCover, wc.Humidity, wc.WindSpeedKmph, wc.WindSpeedMiles,
			wc.WindDir16Point, wc.Pressure, wc.Visibility,
		)
	}

	cmd.Register(w)

	return
}
