// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"regexp"
)

var (
	RegAny     = regexp.MustCompile(`^.+$`)
	RegAlpha   = regexp.MustCompile(`^\w+$`)
	RegHex     = regexp.MustCompile(`^0x[a-fA-F0-9]+$`)
	RegDecimal = regexp.MustCompile(`^[+-]?\d+(\.\d+)?([eE][+-]?\d+)?$`)
	RegOctal   = regexp.MustCompile(`^0[0-7]+$`)
	RegBinary  = regexp.MustCompile(`^0b[01]+$`)
)
