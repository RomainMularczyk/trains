package config

import (
	"trains/src/core/config/resolvers"
	"trains/src/core/config/types"
	"trains/src/core/errors"
	lockFile "trains/src/core/lock/file"
	lockTypes "trains/src/core/lock/types"
)

/*
Resolve the configuration applying variable precedence.
Variable precedence order:
 1. Command line arguments
 2. Environment variables
 3. Configuration file
 4. Default values
*/
func ResolveConfig(path string, provider types.ProviderName) (*types.RuntimeConfig, error) {
	config, err := resolvers.FromFile(path)
	if err != nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: err.Error(),
			Err:     err,
		}
	}

	lockFile, err := lockFile.LoadOrCreate(config.Lock.Path)
	if err != nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidLockError,
			Message: err.Error(),
			Err:     err,
		}
	}

	// Resolve configuration applying variable precedence rules
	resolvers.FromEnv(config)
	resolvers.FromOptions(config)

	runtimeConfig := buildRuntimeConfig(config, provider, lockFile)

	return &runtimeConfig, nil
}

/*
Builds the runtime configuration from configuration files and lock file.
*/
func buildRuntimeConfig(
	config *types.ConfigFile,
	provider types.ProviderName,
	lockFile *lockTypes.LockFile,
) types.RuntimeConfig {
	providerConfig := types.ProviderNameToProviderConfig(provider, *config)

	return types.RuntimeConfig{
		SelectedProvider: types.Provider{
			Name:    provider,
			ApiKey:  providerConfig.ApiKey,
			Model:   providerConfig.Model,
			BaseUrl: providerConfig.BaseUrl,
			Timeout: providerConfig.Timeout,
		},
		Locks: types.CreateLockFileEntries(*lockFile),
	}
}
