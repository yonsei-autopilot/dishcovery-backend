package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func GoogleLogin(ctx context.Context, accessToken string) (*dto.LoginResponse, error) {
	userInfo, err := repository.FetchUserInfo(accessToken)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserById(ctx, userInfo.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user does not exist")
	}

	// TODO - jwt 발급해야 함
	return &dto.LoginResponse{AccessToken: "a", RefreshToken: "r"}, nil
}
