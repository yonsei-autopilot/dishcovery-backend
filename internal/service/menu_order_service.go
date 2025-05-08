package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func OrderMenu(ctx context.Context, id string, request *dto.MenuOrderRequest) (*dto.MenuOrderResponse, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	menuOrder := &answer.MenuOrderAnswer{}
	prompt := getPrompt(user, request)

	_, err = gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithPrompt(prompt).
		ExpectStructuredOutput(menuOrder).
		Generate()
	if err != nil {
		return nil, &fail.FailedExplanationGeneration
	}

	return dto.FromMenuOrderAnswer(menuOrder), nil
}

func getPrompt(user *domain.User, request *dto.MenuOrderRequest) string {
	return fmt.Sprintf(`
Please generate four statements as follows:

1. An order statement in the user's preferred language (%s) for the following menu items. The statement should be clear, polite, and ensure that the disliked foods (%v) are avoided in the order.
2. The same order statement in a foreign language (%s).
3. An inquiry in the user's preferred language asking the waiter or chef if any ingredients in the menu items could potentially include any of the disliked foods (%v). Ensure that the inquiry is polite and respectful, and avoid suggesting the disliked foods are in the dish.
4. The same inquiry statement in the foreign language (%s).

Here are the details of the menu items:
`, user.Language, user.DislikeFoods, request.ForeignLanguage, user.DislikeFoods, request.ForeignLanguage) +
		getMenuItemsDescription(request.Menus) + `
`
}

func getMenuItemsDescription(menus []struct {
	Name        string `json:"name"`
	Count       string `json:"count"`
	Description string `json:"description"`
}) string {
	var menuDescriptions string
	for _, menu := range menus {
		menuDescriptions += fmt.Sprintf("Dish: %s, Quantity: %s, Description: %s\n", menu.Name, menu.Count, menu.Description)
	}
	return menuDescriptions
}
