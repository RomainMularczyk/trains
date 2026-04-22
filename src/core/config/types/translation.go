package configTypes

type Translation struct {
	SourceLanguage Language `validate:"required"`
	TargetLanguage Language
}

type TranslationOverrides struct {
	SourceLanguage *Language
	TargetLanguage *Language
}

/*
Applies the given translation overrides to the translation configuration.
*/
func (t *Translation) Apply(o *TranslationOverrides) {
	if o == nil {
		return
	}
	if o.SourceLanguage != nil {
		t.SourceLanguage = *o.SourceLanguage
	}
	if o.TargetLanguage != nil {
		t.TargetLanguage = *o.TargetLanguage
	}
}
