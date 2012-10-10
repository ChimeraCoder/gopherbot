// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"regexp"
)

// Param represents a single command parameter.
type Param struct {
	Name        string         // Parameter name.
	Description string         // Parameter description.
	Pattern     *regexp.Regexp // Parameter validation pattern.
}
