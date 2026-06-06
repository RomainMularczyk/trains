package testdataEntry

import "trains/src/core/reader/parser"

func TwoEntries() []parser.TranslationEngineEntry {
	return []parser.TranslationEngineEntry{
		{
			Key:    "user.name",
			Target: "My name is {{name}}",
		},
		{
			Key:    "user.age",
			Target: "My age is {{age}}",
		},
	}
}
