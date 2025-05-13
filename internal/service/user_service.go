package service

import (
	"context"

	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/user"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func GetDislikeFoods(ctx context.Context, id string) (*dto.GetDislikeFoodsResponse, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	dislikeFoods := user.DislikeFoods
	if dislikeFoods == nil {
		dislikeFoods = make([]string, 0)
	}

	return dto.FromDislikeFoods(dislikeFoods), nil
}

func UpdateDislikeFoods(ctx context.Context, id string, request *dto.UpdateDislikeFoodsResponse) *fail.Fail {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return &fail.UserNotFound
	}

	user.DislikeFoods = request.DislikeFoods

	err = repository.UpdateUser(ctx, id, user)
	if err != nil {
		return &fail.FailedUpdatingUser
	}

	return nil
}

func GetLanguage(ctx context.Context, id string) (*dto.GetLanguageResponse, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	return dto.FromLanguage(user.Language), nil
}

func UpdateLanguage(ctx context.Context, id string, request *dto.UpdateLanguageRequest) *fail.Fail {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return &fail.UserNotFound
	}

	user.Language = request.Language

	err = repository.UpdateUser(ctx, id, user)
	if err != nil {
		return &fail.FailedUpdatingUser
	}

	return nil
}
