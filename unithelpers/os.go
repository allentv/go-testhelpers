package unithelpers

import "os"

// SetOSEnvVar sets a value and then cleans it up after testing through a 'closer' function.
//
// Since this function sets the value for an environment variable,
// only the closer function is returned.
func (h *TestHelper) SetOSEnvVar(name, value string) func() {
	h.helper()
	os.Setenv(name, value)
	return func() {
		os.Unsetenv(name)
	}
}
