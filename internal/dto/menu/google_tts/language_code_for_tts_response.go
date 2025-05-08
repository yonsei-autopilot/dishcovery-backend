package google_tts

import answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"

type LanguageCodeForGoogleTtsResponse struct {
	LanguageCode string `json:"languageCode"`
}

func FromLanguageCodeForGoogleTtsAnswer(a *answer.LanguageCodeForGoogleTtsAnswer) *LanguageCodeForGoogleTtsResponse {
	return &LanguageCodeForGoogleTtsResponse{LanguageCode: a.LanguageCodeForGoogleTts}
}
