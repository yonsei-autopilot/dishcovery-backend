package answer

type ForeignLanguageOfMenuAnswer struct {
	LanguageName             string `json:"languageName" genai:"description=Language name of the foreign language. it is format of English, Korean, Vietnamese, ...;required"`
	LanguageCodeForGoogleTts string `json:"languageCodeForGoogleTts" genai:"description=Just the language code for the google tts, it should correspond to the foreign language. it is format of the code is like en-US, ko-KR, vi-VN, ...;required"`
}
