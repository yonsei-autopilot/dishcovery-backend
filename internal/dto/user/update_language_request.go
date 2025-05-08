package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type UpdateLanguageRequest struct {
	Language string `json:"language"`
}

func (u *UpdateLanguageRequest) Validate() *fail.Fail {
	if u.Language == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
