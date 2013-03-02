// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package weather

type WeatherData struct {
	Data struct {
		Conditions []struct {
			CloudCover      string `json:"cloudcover"`
			Humidity        string `json:"humidity"`
			ObservationTime string `json:"observation_time"`
			Precip          string `json:"precipMM"`
			Pressure        string `json:"pressure"`
			TempC           string `json:"temp_C"`
			TempF           string `json:"temp_F"`
			Visibility      string `json:"visibility"`
			WeatherCode     string `json:"weatherCode"`
			WeatherDesc     []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			WeatherIconUrl []struct {
				Value string `json:"value"`
			} `json:"weatherIconUrl"`
			WindDir16Point string `json: "winddir16Point"`
			WindDirDegree  string `json: "winddirDegree"`
			WindSpeedKmph  string `json: "windspeedKmph"`
			WindSpeedMiles string `json: "windspeedMiles"`
		} `json:"current_condition"`
		Request []struct {
			Query string
			Type  string
		} `json:"request"`
	} `json:"data"`
}
