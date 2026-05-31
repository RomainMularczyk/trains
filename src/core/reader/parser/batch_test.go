package parser

import (
	"testing"
)

func TestBatchCreateSimple(t *testing.T) {
	units := []TranslationUnit{
		{
			Fullkey:  "user.name",
			Path:     []string{"user", "name"},
			Source:   "Jovan",
			Target:   "",
			Segments: []Segment{{Type: TextSegment, Value: "Jovan"}},
		},
	}

	unitChan := make(chan TranslationUnit, len(units))
	for _, unit := range units {
		unitChan <- unit
	}
	close(unitChan)

	var batches []TranslationBatch
	CreateBatch(unitChan, &batches, 100)

	if len(batches) != 1 {
		t.Errorf("Expected 1 batch, got %d", len(batches))
	}
	if len(batches[0].Units) != 1 {
		t.Errorf("Expected 1 unit, got %d", len(batches[0].Units))
	}
	if batches[0].Units[0].Fullkey != "user.name" {
		t.Errorf("Expected user.name, got %s", batches[0].Units[0].Fullkey)
	}
}
