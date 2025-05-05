package dto

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type SimpleLoginRequest struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
}

func (r *SimpleLoginRequest) Validate() *fail.Fail {
	if r.LoginId == "" || r.Password == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
