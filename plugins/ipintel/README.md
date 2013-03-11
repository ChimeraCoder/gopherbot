## IP-Intel

This plugin provides geolocation information for IP addresses using the
IP Intelligence services at [neustar.biz][1].

[1]: https://ipintelligence.neustar.biz

It requires registration of a free developer account to get access to the
required API keys. These should be placed in the configuration file for
this plugin:

	[api]
	key = xxxxxxxxxxxxxxxxxxxxxxxxxxxx
	shared = xxxxxx

In order for an API request to be successful, a signature is calculated
from the API key and shared secret, plus the current unix timestamp.
The API allows for a 5 minute clock drift. If the system clock this bot is
running on, has more than a 5-minute deviation from the neustar server, the
request will fail with a `403: Not Authorized` error. For this purpose, the
configuration defines the `drift` key, which should hold the number of seconds
to add/subtract to the current timestamp, in order to fall within this error
margin. It should be a positive or negative integer.

	drift = 300


### Commands

* `loc <ipaddress>`: Fetches location data for the given IP address.
  IP Intelligence supports only IPv4 (32-bit) addresses, and they must be in
  standard IPv4 dot-decimal notation. E.g.: `127.0.0.1`, `255.255.255.1`.

### Example

	<steve> ?loc 173.194.65.139
	<bot> steve: 173.194.65.139 (Mapped), Network org.: google inc., Carrier: 
	      google inc., TLD: net, SLD: 1e100. Location: north america/united 
	      states/california/mountain view (37.388020, -122.074310). Postalcode: 
	      94041, Timezone: -8


