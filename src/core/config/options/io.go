package configOptions

import (
	"os"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the IO configuration from environment variables.
*/
func IOFromEnv() (*configTypes.IO, *errors.TrainsError) {
	config := configTypes.IO{}
	if v := os.Getenv("TRAINS_IO_INPUT_FORMAT"); v != "" {
		// TODO: add suggestions of the different formats available
		fileFormat, err := configTypes.FlagToFileFormat(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid IO input format",
				Err:     err,
			}
		}
		config.InputFormat = configTypes.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_OUTPUT_FORMAT"); v != "" {
		fileFormat, err := configTypes.FlagToFileFormat(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid IO output format",
				Err:     err,
			}
		}
		config.OutputFormat = configTypes.FileFormat(fileFormat)
	}
	if v := os.Getenv("TRAINS_IO_SOURCE_PATH"); v != "" {
		config.SourcePath = v
	}
	if v := os.Getenv("TRAINS_IO_TARGET_PATH"); v != "" {
		config.TargetPath = v
	}

	return &config, nil
}
