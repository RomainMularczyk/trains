package testdataEntry

import "trains/src/core/reader/parser"

func OneEntry() []parser.TranslationEngineEntry {
	return []parser.TranslationEngineEntry{
		{
			Key:    "user.name",
			Target: "My name is {{name}}",
		},
	}
}
