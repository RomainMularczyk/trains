package config

import (
	"encoding/json"
	"fmt"
	cmdTypes "trains/src/cli/types"
	configOptions "trains/src/core/config/options"
	"trains/src/core/config/resolvers"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
	lockFile "trains/src/core/lock/file"
	lockTypes "trains/src/core/lock/types"
)

type ConfigLayers struct {
	Default configTypes.ConfigFile
	File    configTypes.ConfigFileOverrides
	Env     configTypes.ConfigFileOverrides
	CLI     configTypes.ConfigFileOverrides
}

func (c ConfigLayers) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf(
			"{\"default\": %s,\"file\": %s,\"env\": %s,\"cli\": %s}",
			c.Default,
			c.File,
			c.Env,
			c.CLI,
		)
	}
	return string(config)
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
	bootstrapLogger *configTypes.StageLogger,
) (*configTypes.RuntimeConfig, *errors.TrainsError) {
	defaultConfig := resolvers.FromDefaults()
	configFilePath := resolveConfigFilePath(cliConfigOptions, defaultConfig)

	fileConfig, err := resolvers.FromFile(configFilePath, bootstrapLogger)
	if err != nil {
		return nil, err
	}

	envConfig, err := resolvers.FromEnv()
	if err != nil {
		return nil, err
	}
	cliConfig, err := resolvers.FromOptions(cliConfigOptions, bootstrapLogger)

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

	runtimeConfig, err := buildRuntimeConfig(mergedConfig, *lockFile)
	if err != nil {
		return nil, err
	}

	return runtimeConfig, nil
}

func MergeConfig(
	configs ConfigLayers,
) configTypes.ConfigFile {
	config := configTypes.ConfigFile{
		Batching:         configs.Default.Batching,
		Config:           configs.Default.Config,
		IO:               configs.Default.IO,
		Lock:             configs.Default.Lock,
		Logging:          configs.Default.Logging,
		Prompt:           configs.Default.Prompt,
		Providers:        configs.Default.Providers,
		SelectedProvider: configs.Default.SelectedProvider,
		Translation:      configs.Default.Translation,
	}

	// File overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.File.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.File.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.File.Lock)
	configOptions.MergeLoggingConfig(&config.Logging, configs.File.Logging)
	configOptions.MergePromptConfig(&config.Prompt, configs.File.Prompt)
	configOptions.MergeProvidersConfig(&config.Providers, configs.File.Providers)
	configOptions.MergeSelectedProviderConfig(
		&config.SelectedProvider,
		configs.File.SelectedProvider,
	)
	configOptions.MergeTranslationConfig(&config.Translation, configs.File.Translation)

	// Env overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.Env.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.Env.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.Env.Lock)
	configOptions.MergeLoggingConfig(&config.Logging, configs.Env.Logging)
	configOptions.MergePromptConfig(&config.Prompt, configs.Env.Prompt)
	configOptions.MergeProvidersConfig(&config.Providers, configs.Env.Providers)
	configOptions.MergeTranslationConfig(&config.Translation, configs.Env.Translation)

	// CLI overrides
	configOptions.MergeBatchingConfig(&config.Batching, configs.CLI.Batching)
	configOptions.MergeIOConfig(&config.IO, configs.CLI.IO)
	configOptions.MergeLockConfig(&config.Lock, configs.CLI.Lock)
	configOptions.MergeLoggingConfig(&config.Logging, configs.CLI.Logging)
	configOptions.MergePromptConfig(&config.Prompt, configs.CLI.Prompt)
	configOptions.MergeProvidersConfig(&config.Providers, configs.CLI.Providers)
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
Resolves the provider configuration from the given config.
*/
func resolveProvider(
	config configTypes.ConfigFile,
) (*configTypes.Provider, *errors.TrainsError) {
	if config.SelectedProvider == "" {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidProviderError,
			Message: "No provider selected. Use --provider/-p to select a provider",
			Err:     fmt.Errorf("No provider selected"),
		}
	}

	if config.Providers[config.SelectedProvider] == nil {
		return nil, &errors.TrainsError{
			Code:    errors.InvalidProviderError,
			Message: "The provider selected does not exist",
			Err:     fmt.Errorf("Invalid provider selected"),
		}
	}

	return config.Providers[config.SelectedProvider], nil
}

/*
Validates the provider configuration.
*/
func validateProvider(
	provider configTypes.Provider,
) *errors.TrainsError {
	if provider.Model == "" {
		return &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Missing provider model",
			Err:     fmt.Errorf("Missing provider model"),
		}
	}
	if provider.BaseUrl == "" {
		return &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Missing provider base URL",
			Err:     fmt.Errorf("Missing provider base URL"),
		}
	}
	if provider.ApiKey == "" {
		return &errors.TrainsError{
			Code:    errors.InvalidConfigError,
			Message: "Missing provider API key",
			Err:     fmt.Errorf("Missing provider API key"),
		}
	}
	return nil
}

/*
Builds the runtime configuration from configuration files and lock file.
*/
func buildRuntimeConfig(
	config configTypes.ConfigFile,
	lockFile lockTypes.LockFile,
) (*configTypes.RuntimeConfig, *errors.TrainsError) {
	provider, err := resolveProvider(config)
	if err != nil {
		return nil, err
	}
	err = validateProvider(*provider)
	if err != nil {
		return nil, err
	}

	logger := configTypes.NewLogger(config.Logging.Level)

	return &configTypes.RuntimeConfig{
		Batching:    config.Batching,
		Config:      config.Config,
		IO:          config.IO,
		Lock:        config.Lock,
		Logger:      logger,
		Prompt:      config.Prompt,
		Provider:    *provider,
		Translation: config.Translation,
		Locks:       lockFile.Entries,
	}, nil
}
