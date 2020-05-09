package unithelpers

import "encoding/json"

// GetJSON returns the JSON representation of the struct
func (h *TestHelper) GetJSON(obj interface{}) string {
	h.Helper()
	data, err := json.Marshal(obj)
	if err != nil {
		h.Fatalf("Cannot convert to JSON. %+v", obj)
	}
	return string(data)
}

// GetObject converts the string data into an object
func (h *TestHelper) GetObject(data string, dataType interface{}) interface{} {
	h.Helper()
	err := json.Unmarshal([]byte(data), &dataType)
	if err != nil {
		h.Fatalf("Cannot convert to object for : %v", data)
	}

	return dataType
}
