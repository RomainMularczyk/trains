package resolvers

import (
	configOptions "trains/src/core/config/options"
	configTypes "trains/src/core/config/types"
)

/*
Resolves the configuration from command line options.
*/
func FromOptions(config *configTypes.ConfigFile) {
	configOptions.ProviderFromOptions(config)
}
