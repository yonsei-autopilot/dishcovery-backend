package dto

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type GoogleLoginRequest struct {
	AccessToken string `json:"accessToken"`
}

func (r *GoogleLoginRequest) Validate() *fail.Fail {
	if r.AccessToken == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
