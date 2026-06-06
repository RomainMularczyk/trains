package utils

import (
	"encoding/json"
	"os"
)

/*
Check if the file is empty.
*/
func IsFileEmpty(fileContent []byte) bool {
	if fileContent == nil {
		return true
	}
	return false
}

/*
Opens a JSON file and unmarshals it into a map.
*/
func OpenJsonFile(path string) any {
	inputFile, _ := os.ReadFile(path)
	var input any
	json.Unmarshal(inputFile, &input)

	return input
}
