// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

const (
	AppName         = "ircb"
	AppVersionMajor = 0
	AppVersionMinor = 1
)

// revision part of the program version.
// This will be set automatically at build time like so:
//
//     go build -ldflags "-X main.AppVersionRev `date -u +%s`"
var AppVersionRev string

func Version() string {
	if len(AppVersionRev) == 0 {
		AppVersionRev = "0"
	}

	return fmt.Sprintf("%s %d.%d.%s (Go runtime %s).\nCopyright (c) 2010-2012, Jim Teeuwen.",
		AppName, AppVersionMajor, AppVersionMinor, AppVersionRev, runtime.Version())
}
