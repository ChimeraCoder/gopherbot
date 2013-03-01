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

