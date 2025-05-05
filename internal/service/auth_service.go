package service

import (
	"context"
	"time"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/token"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/auth"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func GoogleLogin(ctx context.Context, accessToken string) (*dto.LoginResponse, *fail.Fail) {
	userInfo, err := repository.FetchUserInfo(accessToken)
	if err != nil {
		return nil, &fail.UserNotGoogleAuthenticated
	}

	id := getGoogleAuthenticatedUserId(userInfo.Id)
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, addNewUser(ctx, id, userInfo)
	}

	updateLastLogin(ctx, id, user)

	accessToken, refreshToken, fail := token.CreateTokens(id)
	if fail != nil {
		return nil, fail
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func SimpleLogin(ctx context.Context, req *dto.SimpleLoginRequest) (*dto.LoginResponse, *fail.Fail) {
	id := getSimpleAuthenticatedUserId(req.LoginId)

	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	updateLastLogin(ctx, id, user)

	accessToken, refreshToken, fail := token.CreateTokens(id)
	if fail != nil {
		return nil, fail
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func Register(ctx context.Context, req *dto.RegisterRequest) *fail.Fail {
	id := getSimpleAuthenticatedUserId(req.LoginId)

	user := req.ToUser(time.Now())

	err := repository.AddUser(ctx, id, user)
	if err != nil {
		return &fail.FailedSavingUser
	}

	return nil
}

func addNewUser(ctx context.Context, id string, userInfo *dto.UserInfoResponse) *fail.Fail {
	newUser := userInfo.ToUser(time.Now())

	err := repository.AddUser(ctx, id, newUser)
	if err != nil {
		return &fail.FailedSavingUser
	}

	return &fail.UserNotFullyRegistered
}

func updateLastLogin(ctx context.Context, id string, user *domain.User) {
	now := time.Now()
	user.LastLogin = &now

	_ = repository.SetUser(ctx, id, user)
}

func getGoogleAuthenticatedUserId(sub string) string {
	return "google:" + sub
}

func getSimpleAuthenticatedUserId(loginId string) string {
	return "simple:" + loginId
}
