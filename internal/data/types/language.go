package types

// Language represents a language
type Language int

// The differents languages
const (
	LangVO = Language(iota)
	LangVF
	LangVOSTFR
)

// Strings of the languages
const (
	LangVOString     = "VO"
	LangVFString     = "VF"
	LangVOSTFRString = "VOSTFR"
)

// LanguageName maps Language -> String
var LanguageName = map[Language]string{
	LangVO:     LangVOString,
	LangVF:     LangVFString,
	LangVOSTFR: LangVOSTFRString,
}

// LanguageValue maps String -> Language
var LanguageValue = map[string]Language{
	LangVOString:     LangVO,
	LangVFString:     LangVF,
	LangVOSTFRString: LangVOSTFR,
}
