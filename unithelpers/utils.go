package unithelpers

import (
	"fmt"
	"io/ioutil"
)

// LoadData reads the test data file under `testdata` folder
func (h *TestHelper) LoadData(filename string) string {
	h.Helper()
	data, err := ioutil.ReadFile(fmt.Sprintf("testdata/%v", filename))
	if err != nil {
		h.Fatal(err)
	}
	return string(data)
}
