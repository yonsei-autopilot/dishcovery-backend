package dto

import (
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
)

type UserInfoResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserInfoResponse) ToUser(now time.Time) *domain.User {
	return &domain.User{
		Name:         u.Name,
		Language:     nil,
		DislikeFoods: nil,
		AuthProvider: "google",
		CreatedAt:    now,
		LastLogin:    nil,
	}
}
