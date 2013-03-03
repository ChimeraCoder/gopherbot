## Weather

This plugin provides weather report services using the
[worldweatheronline API](http://www.worldweatheronline.com).

It requires registration of a free account to obtain an API key.
This key should be entered in the plugin configuration file.

	[api]
	key = xxxxxxxxxxxxxxxxx

It presents weather data in the following form:

	<bob> ?weather london
	<bot> bob, weather in London, United Kingdom: 3°C/37°F, cloud cover: 0%, 
          humidity: 75%, wind: 4kph/2mph from ENE, pressure: 1026 mb,
          visibility: 10 km


### Commands

* `weather <location>`: Fetches current weather data for the given location.
  The location can be a city or town name.

