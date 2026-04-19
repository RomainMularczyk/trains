package translation

import (
	"encoding/json"
	"math"
	"trains/src/core/parser"
)

type TranslationUnit struct {
	Fullkey  string
	Path     []string
	Source   string
	Target   string
	Segments []parser.Segment
}

func (t TranslationUnit) String() string {
	translationSet, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(translationSet)
}

func (t *TranslationUnit) EstimateTokenNumber() int {
	TOKEN_PER_TEXT_SEGMENT_HEURISTIC := 4.0
	estimateNumberOfTokens := 0
	for _, segment := range t.Segments {
		if segment.Type == parser.TextSegment {
			estimateNumberOfTokens += int(math.Ceil(float64(len(segment.Value)) / TOKEN_PER_TEXT_SEGMENT_HEURISTIC))
		}
	}

	return estimateNumberOfTokens
}

var translationUnitPool chan TranslationUnit = make(chan TranslationUnit, 500)
