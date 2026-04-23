package configOptions

import (
	"os"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the selected provider configuration from environment variables.
*/
func SelectedProviderFromEnv() (*configTypes.ProviderName, *errors.TrainsError) {
	provider := configTypes.ProviderName(os.Getenv("TRAINS_PROVIDER"))

	return &provider, nil
}

/*
Resolves the selected provider configuration from command line options.
*/
func SelectedProviderFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.ProviderName, *errors.TrainsError) {
	return &cliConfigOptions.SelectedProvider, nil
}

/*
Merges the selected provider configuration with the given overrides.
*/
func MergeSelectedProviderConfig(
	dest *configTypes.ProviderName,
	overrides *configTypes.ProviderName,
) {
	dest = overrides
}
