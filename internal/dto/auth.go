package dto

import (
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
)

type GoogleLoginRequest struct {
	AccessToken string `json:"accessToken"`
}

type SimpleLoginRequest struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
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
		Name:         u.Name,
		Language:     nil,
		DislikeFoods: nil,
		CreatedAt:    now,
		LastLogin:    nil,
	}
}
