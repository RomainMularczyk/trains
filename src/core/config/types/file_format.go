package configTypes

import (
	"fmt"
	"trains/src/core/errors"
)

type FileFormat string

const (
	JSON FileFormat = "json"
	YAML FileFormat = "yaml"
)

/*
Converts a flag to a file format.
*/
func FlagToFileFormat(flag string) (FileFormat, *errors.TrainsError) {
	switch flag {
	case "json":
		return JSON, nil
	case "yaml":
		return YAML, nil
	default:
		return JSON, &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Invalid file format",
			Err:     fmt.Errorf("Invalid file format: %s", flag),
		}
	}
}
