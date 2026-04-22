package resolvers

import (
	configOptions "trains/src/core/config/options"
	"trains/src/core/config/types"
)

/*
Resolves the configuration from command line options.
*/
func FromOptions(config *types.ConfigFile) {
	configOptions.ProviderFromOptions(config)
}
