package types

import "fmt"

type FileFormat string

const (
	JSON FileFormat = "json"
	YAML FileFormat = "yaml"
)

func FlagToFileFormat(flag string) (FileFormat, error) {
	switch flag {
	case "json":
		return JSON, nil
	case "yaml":
		return YAML, nil
	default:
		return JSON, fmt.Errorf("Unsupported format: %s", flag)
	}
}
