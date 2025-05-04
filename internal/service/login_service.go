package service

import (
	"context"
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func GoogleLogin(ctx context.Context, accessToken string) (*dto.LoginResponse, error) {
	userInfo, err := repository.FetchUserInfo(accessToken)
	if err != nil {
		return nil, &fail.UserNotGoogleAuthenticated
	}

	_, err = repository.GetUserById(ctx, userInfo.Id)
	if err != nil {
		newUser := userInfo.ToUser(time.Now())

		_, err = repository.AddUser(ctx, newUser)
		if err != nil {
			return nil, &fail.FailedSavingUser
		}

		return nil, &fail.UserNotFullyRegistered
	}

	// TODO - jwt 발급해야 함
	return &dto.LoginResponse{AccessToken: "a", RefreshToken: "r"}, nil
}
