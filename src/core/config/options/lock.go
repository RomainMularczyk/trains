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
Merges the lock configuration with the given overrides.
*/
func MergeLockConfig(dest *configTypes.Lock, overrides *configTypes.LockOverrides) {
	if overrides == nil {
		return
	}
	if overrides.Version != nil {
		dest.Version = *overrides.Version
	}
	if overrides.Path != nil {
		dest.Path = *overrides.Path
	}
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
