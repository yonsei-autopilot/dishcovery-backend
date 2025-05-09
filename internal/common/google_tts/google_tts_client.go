package google_tts

import (
	"context"
	"fmt"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
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

func GetClient() *texttospeech.Client {
	return googleTtsClient
}
