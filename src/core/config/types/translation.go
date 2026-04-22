package configTypes

type Translation struct {
	SourceLanguage Language `validate:"required"`
	TargetLanguage Language
}

type TranslationOverrides struct {
	SourceLanguage *Language
	TargetLanguage *Language
}
