// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package cmd

import (
	"regexp"
	"strconv"
)

// Param represents a single command parameter.
type Param struct {
	Name        string         // Parameter name.
	Description string         // Parameter description.
	Value       string         // Parameter value.
	Pattern     *regexp.Regexp // Parameter validation pattern.
	Optional    bool           // Is this parameter optional?
}

// Valid returns true if the parameter value matches the param pattern.
func (p *Param) Valid() bool {
	if p.Pattern == nil {
		return true
	}

	return p.Pattern.MatchString(p.Value)
}

func (p *Param) S(defaultVal string) string {
	if len(p.Value) > 0 {
		return p.Value
	}
	return defaultVal
}

func (p *Param) B(defaultVal bool) bool {
	v, err := strconv.ParseBool(p.Value)
	if err != nil {
		return defaultVal
	}
	return v
}

func (p *Param) I8(defaultVal int8) int8 {
	v, err := strconv.ParseInt(p.Value, 10, 8)
	if err != nil {
		return defaultVal
	}
	return int8(v)
}

func (p *Param) I16(defaultVal int16) int16 {
	v, err := strconv.ParseInt(p.Value, 10, 16)
	if err != nil {
		return defaultVal
	}
	return int16(v)
}

func (p *Param) I32(defaultVal int32) int32 {
	v, err := strconv.ParseInt(p.Value, 10, 32)
	if err != nil {
		return defaultVal
	}
	return int32(v)
}

func (p *Param) I64(defaultVal int64) int64 {
	v, err := strconv.ParseInt(p.Value, 10, 64)
	if err != nil {
		return defaultVal
	}
	return int64(v)
}

func (p *Param) U8(defaultVal uint8) uint8 {
	v, err := strconv.ParseUint(p.Value, 10, 8)
	if err != nil {
		return defaultVal
	}
	return uint8(v)
}

func (p *Param) U16(defaultVal uint16) uint16 {
	v, err := strconv.ParseUint(p.Value, 10, 16)
	if err != nil {
		return defaultVal
	}
	return uint16(v)
}

func (p *Param) U32(defaultVal uint32) uint32 {
	v, err := strconv.ParseUint(p.Value, 10, 32)
	if err != nil {
		return defaultVal
	}
	return uint32(v)
}

func (p *Param) U64(defaultVal uint64) uint64 {
	v, err := strconv.ParseUint(p.Value, 10, 64)
	if err != nil {
		return defaultVal
	}
	return v
}

func (p *Param) F32(defaultVal float32) float32 {
	v, err := strconv.ParseFloat(p.Value, 32)
	if err != nil {
		return defaultVal
	}
	return float32(v)
}

func (p *Param) F64(defaultVal float64) float64 {
	v, err := strconv.ParseFloat(p.Value, 64)
	if err != nil {
		return defaultVal
	}
	return v
}
