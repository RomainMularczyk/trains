package json

import (
	"testing"
)

func TestReadEmptyJson(t *testing.T) {
	reader := JSONReader{}
	var result map[string]interface{}

	err := reader.Read("../../../../../tests/empty.json", &result)
	if err == nil {
		t.Error("Expected an error if JSON file is empty")
		return
	}
}

func TestReadSimpleJson(t *testing.T) {
	reader := JSONReader{}
	var result map[string]interface{}

	err := reader.Read("../../../../../tests/simple.json", &result)
	if err != nil {
		t.Errorf("Error reading simple.json: %v", err)
	}
}
