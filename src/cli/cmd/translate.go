package cmd

import (
	"trains/src/core/config"
	"trains/src/core/config/resolvers"
	configTypes "trains/src/core/config/types"
	"trains/src/core/orchestration"

	cmdTypes "trains/src/cli/types"

	"github.com/spf13/cobra"
)

var batchingOptions cmdTypes.BatchingOptions
var configOptions cmdTypes.ConfigOptions
var ioOptions cmdTypes.IOOptions
var lockOptions cmdTypes.LockOptions
var promptOptions cmdTypes.PromptOptions
var providerOptions cmdTypes.ProviderOptions
var providersOptions = make(map[configTypes.ProviderName]cmdTypes.ProviderOptions)
var translationOptions cmdTypes.TranslationOptions

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "A translation command",
	Long: `Translates input files from one format to another.
Supports JSON input files with configurable parsing options.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		providerName, err := configTypes.FlagToProvider(providerOptions.Name)
		if err != nil {
			return err
		}
		providersOptions[providerName] = providerOptions

		logger := cmd.Context().Value("logger").(*configTypes.StageLogger)
		logLevel, cmdErr := cmd.Flags().GetString("logLevel")
		if cmdErr != nil {
			return cmdErr
		}

		cliConfigOptions := cmdTypes.CLIConfigOptions{
			Batching: batchingOptions,
			Config:   configOptions,
			IO:       ioOptions,
			Lock:     lockOptions,
			Logging: cmdTypes.LoggingOptions{
				Level: logLevel,
			},
			Prompt: promptOptions,
			Providers: map[configTypes.ProviderName]cmdTypes.ProviderOptions{
				providerName: providerOptions,
			},
			SelectedProvider: providerName,
			Translation:      translationOptions,
		}

		config, err := config.ResolveConfig(cliConfigOptions, logger)
		if err != nil {
			return err
		}

		pipe := orchestration.Pipeline{}
		pipe.Run(*config)
		return nil
	},
}

func init() {
	// --------------------------------------------
	// ----------------- Batching -----------------
	// --------------------------------------------
	translateCmd.Flags().IntVar(
		&batchingOptions.TokenLimit,
		"tokenLimit",
		resolvers.DefaultTokenLimit,
		"Batching token limit",
	)
	translateCmd.Flags().IntVar(
		&batchingOptions.UnitLimit,
		"unitLimit",
		resolvers.DefaultTranslationUnitLimit,
		"Batching translation unit limit",
	)

	// ------------------------------------------
	// ----------------- Config -----------------
	// ------------------------------------------
	translateCmd.Flags().StringVarP(
		&configOptions.Path,
		"config",
		"c",
		resolvers.DefaultConfigPath,
		"Configuration file path",
	)

	// --------------------------------------
	// ----------------- IO -----------------
	// --------------------------------------
	translateCmd.Flags().StringVarP(
		&ioOptions.InputFormat,
		"inputFormat",
		"i",
		string(resolvers.DefaultInputFormat),
		"Input file format",
	)
	translateCmd.Flags().StringVarP(
		&ioOptions.OutputFormat,
		"outputFormat",
		"o",
		string(resolvers.DefaultOutputFormat),
		"Output file format",
	)
	translateCmd.Flags().StringVarP(
		&ioOptions.SourcePath,
		"sourcePath",
		"s",
		resolvers.DefaultSourcePath,
		"Source path",
	)
	translateCmd.Flags().StringVarP(
		&ioOptions.TargetPath,
		"targetPath",
		"t",
		resolvers.DefaultTargetPath,
		"Target path",
	)

	// ----------------------------------------
	// ----------------- Lock -----------------
	// ----------------------------------------
	translateCmd.Flags().StringVarP(
		&lockOptions.Path,
		"lockPath",
		"l",
		resolvers.DefaultLockPath,
		"Lock path",
	)
	translateCmd.Flags().IntVar(
		&lockOptions.Version,
		"lockVersion",
		resolvers.DefaultLockVersion,
		"Lock version",
	)

	// --------------------------------------------
	// ----------------- Provider -----------------
	// --------------------------------------------
	translateCmd.Flags().IntVar(
		&providerOptions.Timeout,
		"timeout",
		resolvers.DefaultProviderTimeout,
		"Provider timeout",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.ApiKey,
		"apiKey",
		"k",
		resolvers.DefaultProviderApiKey,
		"Provider API key",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.Model,
		"model",
		"m",
		resolvers.DefaultProviderModel,
		"Provider model",
	)
	translateCmd.Flags().StringVar(
		&providerOptions.BaseUrl,
		"baseUrl",
		resolvers.DefaultProviderBaseUrl,
		"Provider base URL",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.Name,
		"provider",
		"p",
		string(resolvers.DefaultProviderName),
		"Provider name",
	)

	// ------------------------------------------
	// ----------------- Prompt -----------------
	// ------------------------------------------
	translateCmd.Flags().StringVar(
		&promptOptions.Context,
		"context",
		resolvers.DefaultPromptContext,
		"Prompt context",
	)

	// -----------------------------------------------
	// ----------------- Translation -----------------
	// -----------------------------------------------
	translateCmd.Flags().StringVar(
		&translationOptions.SourceLanguage,
		"srcLang",
		"en",
		"Source language",
	)
	translateCmd.Flags().StringVar(
		&translationOptions.TargetLanguage,
		"targetLang",
		"fr",
		"Target language",
	)

	rootCmd.AddCommand(translateCmd)
}
