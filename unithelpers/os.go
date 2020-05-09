package unithelpers

import "os"

// SetOSEnvVar sets a value and then cleans it up after testing
func (h *TestHelper) SetOSEnvVar(name, value string) func() {
	os.Setenv(name, value)
	return func() {
		defer os.Unsetenv(name)
	}
}
