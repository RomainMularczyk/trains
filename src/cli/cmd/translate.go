/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"trains/src/core/io"

	"github.com/spf13/cobra"
)

var (
	readFormat     string
	baseDir        string
	sourceLanguage string
	targetLanguage string
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "A translation command",
	Long: `Translates input files from one format to another.
Supports JSON input files with configurable parsing options.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFormat, err := io.FlagToFileFormat(readFormat)
		if err != nil {
			fmt.Errorf("Invalid file format: %v", err)
			return
		}

		fmt.Println(sourceLanguage)

		reader, err := io.Reader(fileFormat)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(reader)
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
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// translateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
