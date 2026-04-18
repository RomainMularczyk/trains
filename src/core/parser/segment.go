package parser

type SegmentType int

const (
	TextSegment SegmentType = iota
	PlaceholderSegment
)

type Segment struct {
	Type  SegmentType
	Value string
}

func Segmentize(input string) []Segment {
	placeholders := DetectPlaceholders(input)

	// we first check if the string is empty
	if len(placeholders) == 0 && input == "" {
		return []Segment{}
	}

	if len(placeholders) == 0 {
		return []Segment{{TextSegment, input}}
	}

	segments := make([]Segment, 0)
	previousPlaceholderEndIndex := 0
	// for each placeholder, we check if there is a text segment before it
	for index, placeholder := range placeholders {
		placeholderStartIndex, placeholderEndIndex := placeholder.PatternIndices[0], placeholder.PatternIndices[1]

		// we simply check if a placeholder is starting the segment
		if index == 0 && placeholderStartIndex == 0 {
			segments = append(segments, Segment{PlaceholderSegment, placeholder.Name})
			previousPlaceholderEndIndex = placeholderEndIndex
			continue
		}

		// if the start of the next placeholder is the same as the end of the previous one
		if placeholderStartIndex == previousPlaceholderEndIndex {
			segments = append(segments, Segment{PlaceholderSegment, placeholder.Name})
			previousPlaceholderEndIndex = placeholderEndIndex
			continue
		}

		textSegment := input[previousPlaceholderEndIndex:placeholderStartIndex]
		segments = append(segments, Segment{TextSegment, textSegment}, Segment{PlaceholderSegment, placeholder.Name})
		previousPlaceholderEndIndex = placeholderEndIndex
	}

	// then, we check if there is a text segment after the last placeholder (we also need to check
	if previousPlaceholderEndIndex != len(input) {
		textSegment := input[previousPlaceholderEndIndex:]
		segments = append(segments, Segment{TextSegment, textSegment})
	}

	return segments
}
