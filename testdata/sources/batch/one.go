package testdataBatch

import (
	"trains/src/core/reader/parser"
	testdataUnit "trains/testdata/sources/unit"
)

func OneUnitBatch() parser.TranslationBatch {
	units := testdataUnit.OneUnit()

	var numberOfTokens int
	for _, unit := range units {
		numberOfTokens += unit.EstimateTokenNumber()
	}

	return parser.TranslationBatch{
		Units:       units,
		Size:        numberOfTokens,
		Context:     "",
		ContextSize: 0,
	}
}
