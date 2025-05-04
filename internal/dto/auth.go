package dto

import (
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
)

type LoginRequest struct {
	AccessToken string `json:"accessToken"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserInfoResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserInfoResponse) ToUser(now time.Time) *domain.User {
	return &domain.User{
		Id:           u.Id,
		Name:         u.Name,
		DislikeFoods: "",
		CreatedAt:    now,
		LastLogin:    time.Time{},
	}
}
