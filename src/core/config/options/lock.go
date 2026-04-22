package configOptions

import (
	"os"
	"strconv"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the lock configuration from command line options.
*/
func LockFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.Lock, *errors.TrainsError) {
	config := configTypes.Lock{
		Version: cliConfigOptions.Lock.Version,
		Path:    cliConfigOptions.Lock.Path,
	}
	return &config, nil
}

/*
Resolves the lock configuration from environment variables.
*/
func LockFromEnv() (*configTypes.LockOverrides, *errors.TrainsError) {
	config := configTypes.LockOverrides{}
	if v := os.Getenv("TRAINS_LOCK_VERSION"); v != "" {
		version, err := strconv.Atoi(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid lock version",
				Err:     err,
			}
		}
		config.Version = &version
	}
	if v := os.Getenv("TRAINS_LOCK_PATH"); v != "" {
		config.Path = &v
	}
	return &config, nil
}
