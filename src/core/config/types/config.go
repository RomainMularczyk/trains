package types

import "encoding/json"

type Config struct {
	Provider    Providers
	Translation Translation
	Prompt      Prompt
	Batching    Batching
	IO          IO
}

type CLIOptions struct {
	Provider    *ProvidersOverrides
	Translation *TranslationOverrides
	Prompt      *PromptOverrides
	Batching    *BatchingOverrides
	IO          *IOOverrides
}

/*
Format the configuration as a JSON string.
*/
func (c Config) String() string {
	config, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(config)
}
