package config

import (
	"trains/src/cli/cmd"
	cmdTypes "trains/src/cli/types"
	configOptions "trains/src/core/config/options"
	"trains/src/core/config/resolvers"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
	lockFile "trains/src/core/lock/file"
)

type ConfigLayers struct {
	Default configTypes.ConfigFile
	File    configTypes.ConfigFileOverrides
	Env     configTypes.ConfigFileOverrides
	CLI     configTypes.ConfigFileOverrides
}

/*
Resolve the configuration applying variable precedence.
Variable precedence order:
 1. Command line arguments
 2. Environment variables
 3. Configuration file
 4. Default values
*/
func ResolveConfig(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.RuntimeConfig, error) {
	defaultConfig := resolvers.FromDefaults()
	configFilePath := resolveConfigFilePath(cliConfigOptions, defaultConfig)

	fileConfig, err := resolvers.FromFile(configFilePath)
	if err != nil {
		return nil, err
	}

	envConfig, err := resolvers.FromEnv()
	if err != nil {
		return nil, err
	}
	cliConfig, err := resolvers.FromOptions(cliConfigOptions)

	configs := ConfigLayers{
		Default: defaultConfig,
		File:    *fileConfig,
		Env:     *envConfig,
		CLI:     *cliConfig,
	}

	mergedConfig := MergeConfig(configs)

	lockFile, err := lockFile.LoadOrCreate(mergedConfig.Lock.Path)
	if err != nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidLockError,
			Message: err.Error(),
			Err:     err,
		}
	}

	runtimeConfig := buildRuntimeConfig(mergedConfig)

	return &runtimeConfig, nil
}

func MergeConfig(
	configs ConfigLayers,
) configTypes.ConfigFile {
	config := configTypes.ConfigFile{
		Batching:    configs.Default.Batching,
		Config:      configs.Default.Config,
		IO:          configs.Default.IO,
		Lock:        configs.Default.Lock,
		Prompt:      configs.Default.Prompt,
		Provider:    configs.Default.Provider,
		Translation: configs.Default.Translation,
	}

	// File overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.File.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.File.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.File.Lock)
	configOptions.MergePromptConfig(&config.Prompt, configs.File.Prompt)
	configOptions.MergeProviderConfig(&config.Provider, configs.File.Provider)
	configOptions.MergeTranslationConfig(&config.Translation, configs.File.Translation)

	// Env overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.Env.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.Env.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.Env.Lock)
	configOptions.MergePromptConfig(&config.Prompt, configs.Env.Prompt)
	configOptions.MergeProviderConfig(&config.Provider, configs.Env.Provider)
	configOptions.MergeTranslationConfig(&config.Translation, configs.Env.Translation)

	// CLI overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.CLI.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.CLI.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.CLI.Lock)
	configOptions.MergePromptConfig(&config.Prompt, configs.CLI.Prompt)
	configOptions.MergeProviderConfig(&config.Provider, configs.CLI.Provider)
	configOptions.MergeTranslationConfig(&config.Translation, configs.CLI.Translation)

	return config
}

/*
Resolves the configuration file path between CLI options
and default values.
*/
func resolveConfigFilePath(
	cliConfigOptions cmdTypes.CLIConfigOptions,
	defaultConfig configTypes.ConfigFile,
) string {
	if cliConfigOptions.Config.Path == "" {
		return defaultConfig.Config.Path
	}
	return cliConfigOptions.Config.Path
}

/*
Builds the runtime configuration from configuration files and lock file.
*/
func buildRuntimeConfig(
	config configTypes.ConfigFile,
) configTypes.RuntimeConfig {
	providerConfig := configTypes.ProviderNameToProviderConfig(provider, *config)

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
