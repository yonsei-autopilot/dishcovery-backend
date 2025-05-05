package dto

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (r *RefreshRequest) Validate() *fail.Fail {
	if r.RefreshToken == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
