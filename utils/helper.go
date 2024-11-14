package utils

import (
	"encoding/json"
)

// ToJSON converts an object to JSON string
func ToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// FromJSON converts JSON  to  object
func FromJSON(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}
