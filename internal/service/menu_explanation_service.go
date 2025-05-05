package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func ExplainMenu(ctx context.Context, id string, imageBytes []byte, imageFormat string) (*domain.Menu, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	prompt := createPrompt(user)
	output := &domain.Menu{}

	_, err = gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithImage(imageBytes, imageFormat).
		WithPrompt(prompt).
		ExpectStructuredOutput(output).
		Generate()
	if err != nil {
		return nil, &fail.FailedDescriptionGeneration
	}
	return output, nil
}

func createPrompt(user *domain.User) string {
	return fmt.Sprintf(
		"Given a menu, describe each item in %s. Include the dish name, price, and a detailed explanation of its ingredients, flavors, and characteristics. I dislike %s.",
		user.Language, user.DislikeFoods,
	)
}
