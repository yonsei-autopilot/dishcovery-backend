package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type MenuExplanationRequest struct {
	Name string `json:"name"`
}

func (r *MenuExplanationRequest) Validate() *fail.Fail {
	if r.Name == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}
