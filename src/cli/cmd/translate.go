package cmd

import (
	"fmt"

	"trains/src/core/config"
	"trains/src/core/config/types"
	"trains/src/core/orchestration"

	"github.com/spf13/cobra"
)

var (
	readFormat     string
	baseDir        string
	configPath     string
	lockFilePath   string
	sourceLanguage string
	targetLanguage string
	provider       string
)

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "A translation command",
	Long: `Translates input files from one format to another.
Supports JSON input files with configurable parsing options.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileFormat, err := types.FlagToFileFormat(readFormat)
		provider, err := types.FlagToProvider(provider)
		if err != nil {
			return fmt.Errorf("Invalid provider: %v", err)
		}
		config, err := config.ResolveConfig(configPath, provider)
		if err != nil {
			return fmt.Errorf("Invalid file format: %v", err)
		}

		pipe := orchestration.Pipeline{}
		pipe.Run(baseDir, fileFormat, *config)
		return nil
	},
}

func init() {
	translateCmd.Flags().StringVarP(
		&readFormat,
		"format",
		"f",
		"json",
		"Input file format",
	)
	translateCmd.Flags().StringVarP(
		&baseDir,
		"input",
		"i",
		".",
		"Input directory",
	)
	translateCmd.Flags().StringVarP(
		&sourceLanguage,
		"srcLang",
		"s",
		"en",
		"Source language",
	)
	translateCmd.Flags().StringVarP(
		&targetLanguage,
		"targetLang",
		"t",
		"fr,es",
		"Target language",
	)
	translateCmd.Flags().StringVarP(
		&provider,
		"provider",
		"p",
		"openai",
		"Provider",
	)
	translateCmd.Flags().StringVarP(
		&configPath,
		"config",
		"c",
		"trains.json",
		"Configuration file",
	)
	translateCmd.Flags().StringVarP(
		&lockFilePath,
		"lockFile",
		"l",
		"trains-lock.json",
		"Lock file",
	)
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// translateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
