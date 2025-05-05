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
		Password:     "",
		Name:         u.Name,
		Language:     "",
		DislikeFoods: []string{},
		AuthProvider: "google",
		RefreshToken: "",
		CreatedAt:    now,
		LastLogin:    now,
	}
}
