package answer

type LanguageCodeForGoogleTtsAnswer struct {
	LanguageCodeForGoogleTts string `json:"languageCodeForGoogleTts" genai:"description=Just the language code for the google tts, it should correspond to the foreign language. it is format of the code is like en-US, ko-KR, vi-VN, ...;required"`
}
