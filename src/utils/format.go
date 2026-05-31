package utils

import (
	"encoding/json"
)

/*
Formats the given data as JSON.
*/
func FormatJSON(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}
