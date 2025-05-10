package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type MenuOrderRequest struct {
	ForeignLanguage     string `json:"foreignLanguage"`
	ForeignLanguageCode string `json:"foreignLanguageCode"`
	Menus               []struct {
		Name        string `json:"name"`
		Count       string `json:"count"`
		Description string `json:"description"`
	} `json:"menus"`
}

func (m *MenuOrderRequest) Validate() *fail.Fail {
	if len(m.Menus) == 0 {
		return &fail.RequestValidationFailed
	}
	return nil
}
