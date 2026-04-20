package parser

import (
	"encoding/json"
	"math"
)

type TranslationUnit struct {
	Fullkey  string
	Path     []string
	Source   string
	Target   string
	Segments []Segment
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
		if segment.Type == TextSegment {
			estimateNumberOfTokens += int(math.Ceil(float64(len(segment.Value)) / TOKEN_PER_TEXT_SEGMENT_HEURISTIC))
		}
	}

	return estimateNumberOfTokens
}
