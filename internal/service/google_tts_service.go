package service

import (
	"context"
	"encoding/base64"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type SpeakResponse struct {
	Base64Audio string `json:"base64Audio"`
}

func GetSpeech(ctx context.Context, text string, languageCode string) (string, *fail.Fail) {
	resp, err := google_tts.GetClient().SynthesizeSpeech(ctx,
		&texttospeechpb.SynthesizeSpeechRequest{
			Input: &texttospeechpb.SynthesisInput{
				InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
			},
			Voice: &texttospeechpb.VoiceSelectionParams{
				LanguageCode: languageCode,
				SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
			},
			AudioConfig: &texttospeechpb.AudioConfig{
				AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
			},
		})
	if err != nil {
		return "", &fail.TtsGenerationFailed
	}

	base64Audio := base64.StdEncoding.EncodeToString(resp.AudioContent)

	return base64Audio, nil
}
