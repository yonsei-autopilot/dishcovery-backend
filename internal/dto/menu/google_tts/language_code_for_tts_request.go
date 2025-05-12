package google_tts

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type LanguageCodeForGoogleTtsRequest struct {
	SnippetOfForeignLanguage string `json:"snippetOfForeignLanguage"`
}

func (r *LanguageCodeForGoogleTtsRequest) Validate() *fail.Fail {
	if r.SnippetOfForeignLanguage == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
