package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func GetForeignLanguageOfMenu(ctx context.Context, request *google_tts.LanguageCodeForGoogleTtsRequest) (*google_tts.ForeignLanguageOfMenuResponse, *fail.Fail) {
	foreignLanguageOfMenu := &answer.ForeignLanguageOfMenuAnswer{}
	prompt := createPrompt(request)

	_, err := gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithPrompt(prompt).
		ExpectStructuredOutput(foreignLanguageOfMenu).
		Generate()
	if err != nil {
		return nil, &fail.FailedExplanationGeneration
	}

	return google_tts.FromGeminiAnswer(foreignLanguageOfMenu), nil
}

func createPrompt(request *google_tts.LanguageCodeForGoogleTtsRequest) string {
	return fmt.Sprintf(`You will receive a snippet of text in some foreign language. Identify its language and return a JSON object with exactly two properties:
	{
	"languageName": "<the language name in English>",
	"languageCodeForGoogleTts": "<the Google TTS language code, e.g. en-US, ko-KR, vi-VN>"
	}
	Do not include any other keys or commentary—only valid JSON. Here is the text:
	"%s"`, request.SnippetOfForeignLanguage)
}
