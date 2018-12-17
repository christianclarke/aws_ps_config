package lib

import "os"

// SetEnvironmentVars contained in the given config map
func SetEnvironmentVars(config map[string]string) {
	for key, value := range config {
		os.Setenv(key, value)	
	}
}
