package io

import (
	"testing"
)

func TestDiscoverFiles(t *testing.T) {
	reader, err := Reader(JSON)
	if err != nil {
		t.Errorf("Error creating reader: %v", err)
	}
	files, err := DiscoverFiles("../../../tests", reader)
	if err != nil {
		t.Errorf("Error discovering files: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}
