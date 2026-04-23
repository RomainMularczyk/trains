package configOptions

import (
	"os"
	cmdTypes "trains/src/cli/types"
	configTypes "trains/src/core/config/types"
	"trains/src/core/errors"
)

/*
Resolves the translation configuration from command line options.
*/
func TranslationFromOptions(
	cliConfigOptions cmdTypes.CLIConfigOptions,
) (*configTypes.Translation, *errors.TrainsError) {
	config := configTypes.Translation{}

	if cliConfigOptions.Translation.SourceLanguage != "" {
		srcLang, err := configTypes.GetLanguage(cliConfigOptions.Translation.SourceLanguage)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid source language",
				Err:     err,
			}
		}
		config.SourceLanguage = srcLang
	}

	if cliConfigOptions.Translation.TargetLanguage != "" {
		targetLang, err := configTypes.GetLanguage(cliConfigOptions.Translation.TargetLanguage)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid target language",
				Err:     err,
			}
		}
		config.TargetLanguage = targetLang
	}

	return &config, nil
}

/*
Merges the translation configuration with the given overrides.
*/
func MergeTranslationConfig(
	dest *configTypes.Translation,
	overrides *configTypes.TranslationOverrides,
) {
	if overrides.SourceLanguage != nil {
		dest.SourceLanguage = *overrides.SourceLanguage
	}
	if overrides.TargetLanguage != nil {
		dest.TargetLanguage = *overrides.TargetLanguage
	}
}

/*
Resolves the translation configuration from environment variables.
*/
func TranslationFromEnv() (*configTypes.TranslationOverrides, *errors.TrainsError) {
	config := configTypes.TranslationOverrides{}
	if v := os.Getenv("TRAINS_TRANSLATION_SOURCE_LANGUAGE"); v != "" {
		srcLang, err := configTypes.GetLanguage(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid source language",
				Err:     err,
			}
		}
		config.SourceLanguage = &srcLang
	}
	if v := os.Getenv("TRAINS_TRANSLATION_TARGET_LANGUAGE"); v != "" {
		targetLang, err := configTypes.GetLanguage(v)
		if err != nil {
			return nil, &errors.TrainsError{
				Code:    errors.InvalidConfigError,
				Message: "Invalid target language",
				Err:     err,
			}
		}
		config.TargetLanguage = &targetLang
	}

	return &config, nil
}
