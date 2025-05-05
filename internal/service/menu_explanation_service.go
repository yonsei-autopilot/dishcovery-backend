package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

type Menu struct {
	Items []Item `json:"items" genai:"description=Menu items list;required"`
}

type Item struct {
	Name        string  `json:"name" genai:"description=Item name;required"`
	Description string  `json:"description" genai:"description=Description of the item. Do not include any item name or price info.;required"`
	Price       float32 `json:"price" genai:"description=Price of item"`
}

func ExplainMenu(ctx context.Context, id string, imageBytes []byte, imageFormat string) (string, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return "", &fail.UserNotFound
	}

	prompt := createPrompt(user)
	output := &Menu{}

	result, err := gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithImage(imageBytes, imageFormat).
		WithPrompt(prompt).
		ExpectStructuredOutput(output).
		Generate()
	if err != nil {
		return "", &fail.FailedDescriptionGeneration
	}
	return result, nil
}

func createPrompt(user *domain.User) string {
	lang := "Korean"
	if user.Language != nil {
		lang = *user.Language
	}

	dislikes := "nothing"
	if user.DislikeFoods != nil {
		dislikes = *user.DislikeFoods
	}

	return fmt.Sprintf(
		"Given a menu, describe each item in %s. Include the dish name, price, and a detailed explanation of its ingredients, flavors, and characteristics. I dislike %s.",
		lang, dislikes,
	)
}
