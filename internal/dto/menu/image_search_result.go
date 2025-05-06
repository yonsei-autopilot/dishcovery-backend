package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type ImageSearchResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

func (r *ImageSearchResult) Validate() *fail.Fail {
	if len(r.Items) == 0 {
		return &fail.RequestValidationFailed
	}
	return nil
}
