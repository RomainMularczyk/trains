package configTypes

import (
	"fmt"
	"strings"
)

type Language string

const (
	// AFRICAN LANGUAGES
	Swahili   Language = "sw"
	Zulu      Language = "zu"
	Tsonga    Language = "ts"
	Xhosa     Language = "xh"
	Afrikaans Language = "af"
	Ndebele   Language = "nd"
	Sotho     Language = "st"
	Tswana    Language = "tn"
	Venda     Language = "ve"

	// EUROPEAN LANGUAGES
	English    Language = "en"
	French     Language = "fr"
	Spanish    Language = "es"
	German     Language = "de"
	Dutch      Language = "nl"
	Irish      Language = "ga"
	Italian    Language = "it"
	Portuguese Language = "pt"
	Greek      Language = "el"
	Bulgarian  Language = "bg"
	Polish     Language = "pl"
	Romanian   Language = "ro"
	Swedish    Language = "sv"
	Danish     Language = "da"
	Norwegian  Language = "no"
	Finnish    Language = "fi"
	Hungarian  Language = "hu"
	Czech      Language = "cs"
	Slovenian  Language = "sl"
	Croatian   Language = "hr"
	Serbian    Language = "sr"
	Lithuania  Language = "lt"
	Latvian    Language = "lv"
	Estonian   Language = "et"
	Russian    Language = "ru"

	// MIDDLE EASTERN LANGUAGES
	Turkish Language = "tr"
	Persian Language = "fa"
	Arabic  Language = "ar"

	// ASIAN LANGUAGES
	Japanese   Language = "ja"
	Chinese    Language = "zh"
	Korean     Language = "ko"
	Thai       Language = "th"
	Vietnamese Language = "vi"
	Indonesian Language = "id"
	Filipino   Language = "fil"
)

var languages = map[string]Language{
	"af":  Afrikaans,
	"ar":  Arabic,
	"bg":  Bulgarian,
	"cs":  Czech,
	"da":  Danish,
	"de":  German,
	"el":  Greek,
	"en":  English,
	"es":  Spanish,
	"et":  Estonian,
	"fa":  Persian,
	"fi":  Finnish,
	"fil": Filipino,
	"fr":  French,
	"ga":  Irish,
	"hr":  Croatian,
	"hu":  Hungarian,
	"id":  Indonesian,
	"it":  Italian,
	"ja":  Japanese,
	"ko":  Korean,
	"lt":  Lithuania,
	"lv":  Latvian,
	"nl":  Dutch,
	"no":  Norwegian,
	"pl":  Polish,
	"pt":  Portuguese,
	"ro":  Romanian,
	"ru":  Russian,
	"sl":  Slovenian,
	"sv":  Swedish,
	"sw":  Swahili,
	"th":  Thai,
	"tr":  Turkish,
	"vi":  Vietnamese,
	"zh":  Chinese,
	"zu":  Zulu,
}

func GetLanguage(code string) (Language, error) {
	lang, ok := languages[strings.ToLower(code)]
	if !ok {
		return "", fmt.Errorf("Unsupported language: %s", code)
	}
	return lang, nil
}
