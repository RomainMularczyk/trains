package json

import (
	"testing"
)

func TestParseSimpleJson(t *testing.T) {
	reader := JSONReader{}
	var result map[string]any

	err := reader.Read("../../../../../tests/simple.json", &result)
	if err != nil {
		t.Errorf("Error reading simple.json: %v", err)
	}
	parser := JSONParser{}
	translationSet := parser.Parse(result)

	if len(translationSet.Units) != 6 {
		t.Errorf("Expected 6 units, got %d", len(translationSet.Units))
	}
}
