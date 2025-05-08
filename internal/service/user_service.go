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

	return dto.FromDislikeFoods(user.DislikeFoods), nil
}
