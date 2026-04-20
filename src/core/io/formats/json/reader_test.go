package json

import (
	"testing"
)

func TestReadEmptyJson(t *testing.T) {
	reader := JSONReader{}
	pathChan := make(chan string, 1)
	contentChan := make(chan any, 1)

	pathChan <- "../../../../../tests/empty.json"
	close(pathChan)

	reader.Read(pathChan, contentChan)
	close(contentChan)

	// With current implementation, empty/invalid JSON files are skipped
	// So we just verify no panic occurs
	for range contentChan {
		t.Error("Expected no content for empty JSON file")
	}
}

func TestReadSimpleJson(t *testing.T) {
	reader := JSONReader{}
	pathChan := make(chan string, 1)
	contentChan := make(chan any, 1)

	pathChan <- "../../../../../tests/simple.json"
	close(pathChan)

	reader.Read(pathChan, contentChan)
	close(contentChan)

	result := <-contentChan
	if result == nil {
		t.Error("Expected non-nil result for valid JSON file")
	}
}
