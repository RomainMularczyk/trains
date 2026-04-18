package translation

type TranslationUnit struct {
	Fullkey      string
	Path         []string
	Source       string
	Target       string
	Placeholders []string
	Metadata     map[string]string
}

type TranslationSet struct {
	Units []TranslationUnit
}

var translationUnitPool chan TranslationUnit = make(chan TranslationUnit, 500)
