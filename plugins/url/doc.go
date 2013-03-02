// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

// This plugin detects webpage urls embedded in incoming messages.
// It extracts the urls, fetches their contents from the web and finds the
// page title element. The title is then posted to the channel/user from wence
// the message came.
//
// This plugin has no commands. It simply hooks into PRIVMSG inputs
// and scans them for URLs.
package url
