package cmd

import (
	"fmt"

	"trains/src/core/config/resolvers"
	"trains/src/core/orchestration"

	"github.com/spf13/cobra"
	cmdTypes "trains/src/cli/types"
)

var batchingOptions cmdTypes.BatchingOptions
var configOptions cmdTypes.ConfigOptions
var ioOptions cmdTypes.IOOptions
var lockOptions cmdTypes.LockOptions
var promptOptions cmdTypes.PromptOptions
var providerOptions cmdTypes.ProviderOptions
var translationOptions cmdTypes.TranslationOptions

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "A translation command",
	Long: `Translates input files from one format to another.
Supports JSON input files with configurable parsing options.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cliConfigOptions := cmdTypes.CLIConfigOptions{
			Batching:    batchingOptions,
			Config:      configOptions,
			IO:          ioOptions,
			Lock:        lockOptions,
			Prompt:      promptOptions,
			Provider:    providerOptions,
			Translation: translationOptions,
		}
		config, err := config.ResolveConfig(cliConfigOptions)
		if err != nil {
			return fmt.Errorf("Invalid file format: %v", err)
		}

		pipe := orchestration.Pipeline{}
		pipe.Run(config)
		return nil
	},
}

func init() {
	// --------------------------------------------
	// ----------------- Batching -----------------
	// --------------------------------------------
	translateCmd.Flags().IntVarP(
		&batchingOptions.TokenLimit,
		"tokenLimit",
		"bt",
		resolvers.DefaultTokenLimit,
		"Batching token limit",
	)
	translateCmd.Flags().IntVarP(
		&batchingOptions.UnitLimit,
		"unitLimit",
		"bu",
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
	translateCmd.Flags().IntVarP(
		&lockOptions.Version,
		"lockVersion",
		"ve",
		resolvers.DefaultLockVersion,
		"Lock version",
	)

	// --------------------------------------------
	// ----------------- Provider -----------------
	// --------------------------------------------
	translateCmd.Flags().IntVarP(
		&providerOptions.Timeout,
		"timeout",
		"to",
		resolvers.DefaultProviderTimeout,
		"Provider timeout",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.ApiKey,
		"apiKey",
		"pk",
		resolvers.DefaultProviderApiKey,
		"Provider API key",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.Model,
		"model",
		"pm",
		resolvers.DefaultProviderModel,
		"Provider model",
	)
	translateCmd.Flags().StringVarP(
		&providerOptions.BaseUrl,
		"baseUrl",
		"pu",
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
	translateCmd.Flags().StringVarP(
		&promptOptions.Context,
		"context",
		"ctx",
		resolvers.DefaultPromptContext,
		"Prompt context",
	)

	// -----------------------------------------------
	// ----------------- Translation -----------------
	// -----------------------------------------------
	translateCmd.Flags().StringVarP(
		&translationOptions.SourceLanguage,
		"srcLang",
		"sl",
		"en",
		"Source language",
	)
	translateCmd.Flags().StringVarP(
		&translationOptions.TargetLanguage,
		"targetLang",
		"tl",
		"fr",
		"Target language",
	)

	rootCmd.AddCommand(translateCmd)
}
