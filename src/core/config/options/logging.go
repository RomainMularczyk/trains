package configOptions

import (
	"os"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolve the logging configuration from command line options.
*/
func LoggingFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.Logging, *errors.TrainsError) {
	loggingLevel, err := ResolveLoggingLevel(cliConfigOptions.Logging.Level)
	if err != nil {
		return nil, err
	}
	config := configTypes.Logging{
		Level: *loggingLevel,
	}

	return &config, nil
}

/*
Resolves the logging configuration from environment variables.
*/
func LoggingFromEnv() (*configTypes.LoggingOverrides, *errors.TrainsError) {
	config := configTypes.LoggingOverrides{}
	if v := os.Getenv("TRAINS_LOGGING_LEVEL"); v != "" {
		level, err := ResolveLoggingLevel(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid logging level",
				Err:     err,
			}
		}
		config.Level = level
	}
	return &config, nil
}

/*
Merges the logging configuration with the given overrides.
*/
func MergeLoggingConfig(dest *configTypes.Logging, overrides *configTypes.LoggingOverrides) {
	if overrides == nil {
		return
	}
	if overrides.Level != nil {
		dest.Level = *overrides.Level
	}
}

/*
Resolves the logging level from the given string.
*/
func ResolveLoggingLevel(level string) (
	*configTypes.LoggingLevel,
	*errors.TrainsError,
) {
	var result configTypes.LoggingLevel
	switch level {
	case "error":
		result = configTypes.Error
	case "warn":
		result = configTypes.Warn
	case "info":
		result = configTypes.Info
	case "debug":
		result = configTypes.Debug
	default:
		result = configTypes.Error
	}
	return &result, nil
}
