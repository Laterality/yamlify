package main

import (
	"encoding/json"

	"github.com/go-yaml/yaml"
)

func ConvertFromJsonString(jsonBytes *[]byte) ([]byte, error) {
	var unmarshaled interface{}
	json.Unmarshal(*jsonBytes, &unmarshaled)

	yamlBytes, err := yaml.Marshal(unmarshaled)

	return yamlBytes, err
}
