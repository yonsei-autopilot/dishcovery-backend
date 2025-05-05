package service

import (
	"context"
	"log"
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

	id := getGoogleAuthenticatedUserId(userInfo)
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, addNewUser(ctx, id, userInfo)
	}

	log.Printf("%s", user.Name)

	// TODO - jwt 발급해야 함
	return &dto.LoginResponse{AccessToken: "a", RefreshToken: "r"}, nil
}

func SimpleLogin(ctx context.Context, req *dto.SimpleLoginRequest) (*dto.LoginResponse, error) {
	id := getSimpleAuthenticatedUserId(req)

	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotRegistered
	}

	log.Printf("%s", user.Name)

	// TODO - jwt 발급해야 함
	return &dto.LoginResponse{AccessToken: "a", RefreshToken: "r"}, nil
}

func addNewUser(ctx context.Context, id string, userInfo *dto.UserInfoResponse) error {
	newUser := userInfo.ToUser(time.Now())

	err := repository.AddUser(ctx, id, newUser)
	if err != nil {
		return &fail.FailedSavingUser
	}

	return &fail.UserNotFullyRegistered
}

func getGoogleAuthenticatedUserId(userInfo *dto.UserInfoResponse) string {
	return "google:" + userInfo.Id
}

func getSimpleAuthenticatedUserId(request *dto.SimpleLoginRequest) string {
	return "simple:" + request.LoginId
}
