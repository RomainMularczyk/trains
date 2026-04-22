package configOptions

import (
	"os"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the IO configuration from command line options.
*/
func IOFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.IO, *errors.TrainsError) {
	config := configTypes.IO{}

	if cliConfigOptions.IO.InputFormat != "" {
		fileFormat, err := configTypes.FlagToFileFormat(cliConfigOptions.IO.InputFormat)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid IO input format",
				Err:     err,
			}
		}
		config.InputFormat = fileFormat
	}

	if cliConfigOptions.IO.OutputFormat != "" {
		fileFormat, err := configTypes.FlagToFileFormat(cliConfigOptions.IO.OutputFormat)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid IO output format",
				Err:     err,
			}
		}
		config.OutputFormat = fileFormat
	}

	config.SourcePath = cliConfigOptions.IO.SourcePath
	config.TargetPath = cliConfigOptions.IO.TargetPath

	return &config, nil
}

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
