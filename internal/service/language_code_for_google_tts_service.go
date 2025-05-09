package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func GetLanguageCodeForGoogleTts(ctx context.Context, request *google_tts.LanguageCodeForGoogleTtsRequest) (*google_tts.LanguageCodeForGoogleTtsResponse, *fail.Fail) {
	languageCode := &answer.LanguageCodeForGoogleTtsAnswer{}
	prompt := fmt.Sprintf("Just the language code for the google tts, it should correspond to the foreign language (%s). it is format of the code is like en-US, ko-KR, vi-VN, ...", request.Language)

	_, err := gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithPrompt(prompt).
		ExpectStructuredOutput(languageCode).
		Generate()
	if err != nil {
		return nil, &fail.FailedExplanationGeneration
	}

	return google_tts.FromLanguageCodeForGoogleTtsAnswer(languageCode), nil
}
