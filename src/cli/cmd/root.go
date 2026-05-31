/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	cmdTypes "trains/src/cli/types"
	configOpts "trains/src/core/config/options"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"

	"github.com/spf13/cobra"
)

var verbose bool
var loggingOptions cmdTypes.LoggingOptions

var rootCmd = &cobra.Command{
	Use:   "trains",
	Short: "A brief description of your application",
	Long: `A CLI tool for batch translating files between various formats.
Supports JSON format with customizable parsing and writing options.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		level, err := configOpts.ResolveLoggingLevel(loggingOptions.Level)
		if err != nil {
			return err
		}
		bootstrapLogger := configTypes.NewLogger(*level)

		ctx := context.WithValue(cmd.Context(), "logger", bootstrapLogger)
		cmd.SetContext(ctx)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		errorMsg := errors.FormatError(err, true)
		fmt.Println(errorMsg)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trains.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(
		&loggingOptions.Level,
		"logLevel",
		string(configTypes.Info),
		"Logging level",
	)

	rootCmd.PersistentFlags().BoolVar(
		&verbose,
		"verbose",
		false,
		"Enable verbose logging",
	)
}
