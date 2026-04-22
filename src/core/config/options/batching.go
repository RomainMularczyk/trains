package configOptions

import (
	"os"
	"strconv"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the batching configuration from command line options.
*/
func BatchingFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.Batching, *errors.TrainsError) {
	config := configTypes.Batching{
		TokenLimit: cliConfigOptions.Batching.TokenLimit,
		UnitLimit:  cliConfigOptions.Batching.UnitLimit,
	}
	return &config, nil
}

/*
Resolves the batching configuration from environment variables.
*/
func BatchingFromEnv() (*configTypes.Batching, *errors.TrainsError) {
	config := configTypes.Batching{}
	if v := os.Getenv("TRAINS_BATCHING_TOKEN_LIMIT"); v != "" {
		tokenLimit, err := strconv.Atoi(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid batching token limit",
				Err:     err,
			}
		}
		config.TokenLimit = tokenLimit
	}

	if v := os.Getenv("TRAINS_BATCHING_UNIT_LIMIT"); v != "" {
		unitLimit, err := strconv.Atoi(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid batching unit limit",
				Err:     err,
			}
		}
		config.UnitLimit = unitLimit
	}

	return &config, nil
}
