package unithelpers

import (
	"io/ioutil"
	"path/filepath"
)

// LoadData reads the test data file under `testdata` folder
func (h *TestHelper) LoadData(filename string) string {
	h.helper()
	data, err := ioutil.ReadFile(filepath.Join(h.TestDataFolder, filename))
	if err != nil {
		h.fatal(err)
	}
	return string(data)
}
