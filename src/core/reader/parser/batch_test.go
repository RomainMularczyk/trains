package parser

import (
	"testing"
	configTypes "trains/src/core/config/types"
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

	batches := make(chan TranslationBatch, 0)
	for _, unit := range units {
		batches <- TranslationBatch{Units: []TranslationUnit{unit}}
	}

	if len(batches) != 1 {
		t.Errorf("Expected 1 batch, got %d", len(batches))
	}

	CreateBatch(unitChan, batches, configTypes.RuntimeConfig{Batching: configTypes.Batching{TokenLimit: 100}})
}
