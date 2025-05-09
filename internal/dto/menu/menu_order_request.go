package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type GetMenuOrderTextsRequest struct {
	ForeignLanguage string `json:"foreignLanguage"`
	Menus           []struct {
		Name        string `json:"name"`
		Count       string `json:"count"`
		Description string `json:"description"`
	} `json:"menus"`
}

func (m *GetMenuOrderTextsRequest) Validate() *fail.Fail {
	if len(m.Menus) == 0 {
		return &fail.RequestValidationFailed
	}
	return nil
}
