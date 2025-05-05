package dto

import (
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type RegisterRequest struct {
	LoginId      string `json:"loginId"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Language     string `json:"language"`
	DislikeFoods string `json:"dislikeFoods"`
}

func (r *RegisterRequest) Validate() *fail.Fail {
	if r.LoginId == "" || r.Password == "" || r.Name == "" || r.Language == "" || r.DislikeFoods == "" {
		return &fail.RequestValidationFailed
	}
	return nil
}

func (r *RegisterRequest) ToUser(now time.Time) *domain.User {
	return &domain.User{
		Name:         r.Name,
		Language:     &r.Language,
		DislikeFoods: &r.DislikeFoods,
		AuthProvider: "simple",
		CreatedAt:    now,
		LastLogin:    nil,
	}
}
