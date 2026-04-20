package configOptions

import (
	"os"
	"strconv"
	"trains/src/core/config/types"
)

func ProviderFromEnv(config *types.Config) {
	OpenAIFromEnv(config)
	AnthropicFromEnv(config)
	GoogleFromEnv(config)
	MistralFromEnv(config)
}

func OpenAIFromEnv(config *types.Config) {
	if v := os.Getenv("TRAINS_PROVIDER_OPENAI_API_KEY"); v != "" {
		config.Provider.OpenAI.ApiKey = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_OPENAI_MODEL"); v != "" {
		config.Provider.OpenAI.Model = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_OPENAI_BASE_URL"); v != "" {
		config.Provider.OpenAI.BaseUrl = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_OPENAI_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			config.Provider.OpenAI.Timeout = timeout
		}
	}
}

func AnthropicFromEnv(config *types.Config) {
	if v := os.Getenv("TRAINS_PROVIDER_ANTHROPIC_API_KEY"); v != "" {
		config.Provider.Anthropic.ApiKey = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_ANTHROPIC_MODEL"); v != "" {
		config.Provider.Anthropic.Model = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_ANTHROPIC_BASE_URL"); v != "" {
		config.Provider.Anthropic.BaseUrl = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_ANTHROPIC_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			config.Provider.Anthropic.Timeout = timeout
		}
	}
}

func GoogleFromEnv(config *types.Config) {
	if v := os.Getenv("TRAINS_PROVIDER_GOOGLE_API_KEY"); v != "" {
		config.Provider.Google.ApiKey = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_GOOGLE_MODEL"); v != "" {
		config.Provider.Google.Model = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_GOOGLE_BASE_URL"); v != "" {
		config.Provider.Google.BaseUrl = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_GOOGLE_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			config.Provider.Google.Timeout = timeout
		}
	}
}

func MistralFromEnv(config *types.Config) {
	if v := os.Getenv("TRAINS_PROVIDER_MISTRAL_API_KEY"); v != "" {
		config.Provider.Mistral.ApiKey = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_MISTRAL_MODEL"); v != "" {
		config.Provider.Mistral.Model = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_MISTRAL_BASE_URL"); v != "" {
		config.Provider.Mistral.BaseUrl = v
	}
	if v := os.Getenv("TRAINS_PROVIDER_MISTRAL_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			config.Provider.Mistral.Timeout = timeout
		}
	}
}
