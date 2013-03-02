// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package plugin

import (
	"github.com/jteeuwen/ini"
	"path/filepath"
)

// ConfigPath returns the fully qualified path for the
// given plugin's configuration file.
func ConfigPath(profile, name string) string {
	path := filepath.Join(profile, "plugins")
	path = filepath.Join(path, name)
	return filepath.Join(path, "config.ini")
}

// LoadConfig reads the ini configuration file for the given plugin.
// Returns nil if the file does not exist.
func LoadConfig(profile, name string) *ini.File {
	ini := ini.New()
	err := ini.Load(ConfigPath(profile, name))

	if err != nil {
		return nil
	}

	return ini
}
