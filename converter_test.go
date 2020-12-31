package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFromJSON(t *testing.T) {
	jsonFile, _ := os.Open("./test.json")
	yamlFile, _ := os.Open("./test.yaml")

	yamlBytes, _ := ioutil.ReadAll(yamlFile)

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	convertedYamlStr, _ := ConvertFromJsonString(&jsonBytes)

	if string(yamlBytes) != string(convertedYamlStr) {
		t.Errorf("Result not matched\nExpected:\n%s\n\nActual:\n%s\n", string(yamlBytes), string(convertedYamlStr))
	}
}
