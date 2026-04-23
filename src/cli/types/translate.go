package cmdTypes

import configTypes "trains/src/core/config/types"

type BatchingOptions struct {
	TokenLimit int
	UnitLimit  int
}

type ConfigOptions struct {
	Path string
}

type IOOptions struct {
	InputFormat  string
	OutputFormat string
	SourcePath   string
	TargetPath   string
}

type LockOptions struct {
	Version int
	Path    string
}

type PromptOptions struct {
	Context string
}

type ProvidersOptions map[configTypes.ProviderName]ProviderOptions

type ProviderOptions struct {
	Name    string
	ApiKey  string
	Model   string
	BaseUrl string
	Timeout int
}

type TranslationOptions struct {
	SourceLanguage string
	TargetLanguage string
}

type CLIConfigOptions struct {
	Batching    BatchingOptions
	Config      ConfigOptions
	IO          IOOptions
	Lock        LockOptions
	Prompt      PromptOptions
	Providers   ProvidersOptions
	Translation TranslationOptions
}
