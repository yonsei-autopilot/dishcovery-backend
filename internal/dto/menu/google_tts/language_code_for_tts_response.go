package google_tts

import answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"

type ForeignLanguageOfMenuResponse struct {
	LanguageName string `json:"languageName"`
	LanguageCode string `json:"languageCode"`
}

func FromGeminiAnswer(a *answer.ForeignLanguageOfMenuAnswer) *ForeignLanguageOfMenuResponse {
	return &ForeignLanguageOfMenuResponse{
		LanguageName: a.LanguageName,
		LanguageCode: a.LanguageCodeForGoogleTts,
	}
}
