package unithelpers

import "encoding/json"

// GetJSON returns the JSON representation of the struct
func (h *TestHelper) GetJSON(obj interface{}) string {
	h.helper()
	data, err := json.Marshal(obj)
	if err != nil {
		h.fatalf("Cannot convert to JSON. %+v", obj)
	}
	return string(data)
}
