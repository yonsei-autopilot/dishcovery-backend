package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/google_search"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func ExplainMenu(ctx context.Context, id string, request *dto.MenuExplanationRequest) (*dto.MenuExplanationResponse, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		return nil, &fail.UserNotFound
	}

	imageSearchResult, failure := google_search.SearchMenuImage(request)
	if failure != nil {
		return nil, failure
	}

	menuExplanation := &domain.MenuExplanation{}
	prompt := makePrompt(user, request)

	_, err = gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithPrompt(prompt).
		ExpectStructuredOutput(menuExplanation).
		Generate()
	if err != nil {
		return nil, &fail.FailedExplanationGeneration
	}

	return dto.NewMenuExplanationResponse(imageSearchResult, menuExplanation), nil
}

func makePrompt(user *domain.User, request *dto.MenuExplanationRequest) string {
	return fmt.Sprintf(`
You are given a menu item name: **%s**

Your task is to:
1. Provide a general **description** of the dish in %s.
2. List the **typical ingredients** that are usually included.
3. Based on the user's preferences, highlight any ingredients to be **careful about**.
   - The user dislikes: %s

Be detailed but concise. Output **only** a structured JSON object matching the following format:
{
  "name": "...",
  "description": "...",
  "ingredients": "...",
  "whatToBeCareful": "..."
}
`, request.Name, user.Language, user.DislikeFoods)
}
