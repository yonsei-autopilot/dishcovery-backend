package google_tts

import (
	"context"
	"fmt"

	"encoding/base64"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"google.golang.org/api/option"
)

var googleTtsClient *texttospeech.Client

func InitializeGoogleTtsClient() {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile("google-tts-application-credentials.json"))
	if err != nil {
		panic(fmt.Sprintf("Failed to create Google Tts client: %v", err))
	}

	googleTtsClient = client
}

func GetSpeech(ctx context.Context, text string, languageCode string) (string, *fail.Fail) {
	resp, err := googleTtsClient.SynthesizeSpeech(ctx,
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
