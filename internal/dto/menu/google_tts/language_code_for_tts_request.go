package google_tts

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type LanguageCodeForGoogleTtsRequest struct {
	Language string `json:"language"`
}

func (r *LanguageCodeForGoogleTtsRequest) Validate() *fail.Fail {
	if r.Language == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
