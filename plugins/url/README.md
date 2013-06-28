## URL

This plugin detects webpage urls embedded in incoming messages.
It extracts the urls, fetches their contents from the web and finds the
page title element. The title is then posted to the channel/user from wence
the message came.

For example:

    <someuser> http://www.youtube.com/watch?v=dQw4w9WgXcQ
    <bot> someuser's link shows: Rick Astley - Never Gonna Give You Up - YouTube

It's configuration file can present a list of regular expression patterns.
These patterns represent (partial) urls, which should be excluded from the
lookup. For example, to ignore all links to imgur.com and those ending
in file extensions for images and video:

	[exclude]
	url < \.imgur\.com
	url < \.(jpe?g|png|gif|bmp|tga|tiff)$
	url < \.(mpe?g|avi|mp[1-5])$


This plugin is capable of recognizing links to specific tweets, and fetches the full tweet text accordingly. For example:

````
             │17:02:43   +chimeracoder | https://twitter.com/chimeracoder/status/334355537095426048       │ 
             │17:02:44        gophrbot | chimeracoder's tweet shows: http://t.co/VfVXPJAefq is finally    │
             │                         | updated! Woohoo! #golang                                         │
````

To enable this functionality, you must set the following environment values to valid values:

````sh
export TWITTER_CONSUMER_KEY=""
export TWITTER_CONSUMER_SECRET=""

export TWITTER_ACCESS_TOKEN=""
export TWITTER_ACCESS_TOKEN_SECRET=""
````

If the values are missing or invalid, or if fetching the tweet text fails for any reason, the plugin fails gracefully by falling back on fetching the <title> attribute

