package testdataUnit

import "trains/src/core/reader/parser"

func SimpleTranslationUnits() []parser.TranslationUnit {
	return []parser.TranslationUnit{
		{
			Fullkey: "user.custom_greetings",
			Path:    []string{"user", "custom_greetings"},
			Source:  "hi {{name}}, welcome to the team!",
			Target:  "",
			Segments: []parser.Segment{
				{Type: parser.TextSegment, Value: "hi "},
				{
					Type:  parser.PlaceholderSegment,
					Value: "name",
					Placeholder: &parser.Placeholder{
						Name:           "name",
						NameIndices:    []int{5, 9},
						Pattern:        "{{name}}",
						PatternIndices: []int{3, 11},
					},
				},
				{Type: parser.TextSegment, Value: ", welcome to the team!"},
			},
		},
		{
			Fullkey: "user.custom_bye",
			Path:    []string{"user", "custom_bye"},
			Source:  "see you later, {{name}}!",
			Target:  "",
			Segments: []parser.Segment{
				{Type: parser.TextSegment, Value: "see you later, "},
				{
					Type:  parser.PlaceholderSegment,
					Value: "name",
					Placeholder: &parser.Placeholder{
						Name:           "name",
						NameIndices:    []int{17, 21},
						Pattern:        "{{name}}",
						PatternIndices: []int{15, 23},
					},
				},
				{Type: parser.TextSegment, Value: "!"},
			},
		},
		{
			Fullkey: "grade.description",
			Path:    []string{"grade", "description"},
			Source:  "you achieved a grade of {{grade}}",
			Target:  "",
			Segments: []parser.Segment{
				{Type: parser.TextSegment, Value: "you achieved a grade of "},
				{
					Type:  parser.PlaceholderSegment,
					Value: "grade",
					Placeholder: &parser.Placeholder{
						Name:           "grade",
						NameIndices:    []int{26, 31},
						Pattern:        "{{grade}}",
						PatternIndices: []int{24, 33},
					},
				},
			},
		},
		{
			Fullkey:  "steps.0",
			Path:     []string{"steps", "0"},
			Source:   "one",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "one"}},
		},
		{
			Fullkey:  "steps.1",
			Path:     []string{"steps", "1"},
			Source:   "two",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "two"}},
		},
		{
			Fullkey:  "steps.2",
			Path:     []string{"steps", "2"},
			Source:   "three",
			Target:   "",
			Segments: []parser.Segment{{Type: parser.TextSegment, Value: "three"}},
		},
	}
}
