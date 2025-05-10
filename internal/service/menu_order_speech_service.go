package service

import (
	"context"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/google_tts"
	dto_google_tts "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func GetMenuOrderSpeech(ctx context.Context, request *dto_google_tts.MenuOrderSpeechRequest) (*dto_google_tts.MenuOrderSpeechResponse, *fail.Fail) {
	base64Audio, fail := google_tts.GetSpeech(ctx, request.MenuOrderText, request.LanguageCode)
	if fail != nil {
		return nil, fail
	}

	return &dto_google_tts.MenuOrderSpeechResponse{Base64Audio: base64Audio}, nil
}
