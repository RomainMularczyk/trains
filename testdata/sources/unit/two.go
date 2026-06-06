package testdataUnit

import "trains/src/core/reader/parser"

func TwoUnits() []parser.TranslationUnit {
	return []parser.TranslationUnit{
		{
			Fullkey: "user.name",
			Path:    []string{"user", "name"},
			Source:  "My name is {{name}}",
			Target:  "",
			Segments: []parser.Segment{
				parser.Segment{
					Type:        parser.TextSegment,
					Value:       "My name is ",
					Placeholder: nil,
				},
				parser.Segment{
					Type:  parser.PlaceholderSegment,
					Value: "name",
					Placeholder: &parser.Placeholder{
						Name:           "name",
						NameIndices:    []int{19, 24},
						Pattern:        "{{name}}",
						PatternIndices: []int{17, 25}},
				},
			},
		},
		{
			Fullkey: "user.age",
			Path:    []string{"user", "age"},
			Source:  "My age is {{age}}",
			Target:  "",
			Segments: []parser.Segment{
				parser.Segment{
					Type:        parser.TextSegment,
					Value:       "My age is ",
					Placeholder: nil,
				},
				parser.Segment{
					Type:  parser.PlaceholderSegment,
					Value: "age",
					Placeholder: &parser.Placeholder{
						Name:           "age",
						NameIndices:    []int{19, 24},
						Pattern:        "{{age}}",
						PatternIndices: []int{17, 25}},
				},
			},
		},
	}
}
