package config

type Provider struct {
	ApiKey  string `validate:"required"`
	Model   string `validate:"required"`
	BaseUrl string `validate:"required"`
	Timeout int    `validate:"required"`
}

type ProviderOverrides struct {
	ApiKey  *string
	Model   *string
	BaseUrl *string
	Timeout *int
}
