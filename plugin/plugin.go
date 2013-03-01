// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package plugin

import (
	"path/filepath"
)

// ConfigPath returns the fully qualified path for the
// given plugin's configuration file.
func ConfigPath(profile, name string) string {
	path := filepath.Join(profile, "plugins")
	path = filepath.Join(path, name)
	return filepath.Join(path, "config.ini")
}
