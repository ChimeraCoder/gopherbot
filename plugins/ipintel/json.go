// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ipintel

// Response holds the unmarshalled response data for a single API call.
// Fields not available in the free api subscription hhave been omitted.
type Response struct {
	IPInfo struct {
		IPAddress string `json:"ip_address"`
		IPType    string `json:"ip_type"`

		Network struct {
			Organization   string `json:"organization"`
			Carrier        string `json:"carrier"`
			ASN            int    `json:"asn"`
			ConnectionType string `json:"connection_type"`
			LineSpeed      string `json:"line_speed"`
			RoutingType    string `json:"ip_routing_type"`

			Domain struct {
				TLD string `json:"tld"`
				SLD string `json:"sld"`
			} `json:"Domain"`
		} `json:"Network"`

		Location struct {
			Continent string  `json:"continent"`
			Latitude  float32 `json:"latitude"`
			Longitude float32 `json:"longitude"`

			CountryData struct {
				Name string `json:"country"`
				Code string `json:"country_code""`
				CF   int    `json:"country_cf"`
			} `json:"CountryData"`

			StateData struct {
				Name string `json:"state"`
				CF   int    `json:"state_cf"`
			} `json:"StateData"`

			CityData struct {
				Name       string `json:"city"`
				PostalCode string `json:"postal_code"`
				TimeZone   int    `json:"time_zone"`
				CF         int    `json:"city_cf"`
			} `json:"StateData"`
		} `json:"location"`
	} `json:"ipinfo"`
}
