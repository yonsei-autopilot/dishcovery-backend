package google_tts

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type MenuOrderSpeechRequest struct {
	MenuOrderText string `json:"menuOrderText"`
	LanguageCode  string `json:"languageCode"`
}

func (r *MenuOrderSpeechRequest) Validate() *fail.Fail {
	if r.MenuOrderText == "" || r.LanguageCode == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
