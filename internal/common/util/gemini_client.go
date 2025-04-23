package util

import (
	"context"
	"errors"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AskGemini(imageBytes []byte, imageFormat string, prompt string) (string, error) {
	ctx := context.Background()

	apiKey, err := GetEnv("GEMINI_API_KEY")
	if err != nil {
		return "", errors.New("missing Gemini API key")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", errors.New("failed to create Gemini client")
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	resp, err := model.GenerateContent(ctx,
		genai.ImageData(imageFormat, imageBytes),
		genai.Text(prompt),
	)
	if err != nil {
		return "", errors.New("Gemini generation failed")
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", errors.New("Gemini returned empty result")
	}

	return extractTextFromContent(resp.Candidates[0].Content), nil
}

func extractTextFromContent(content *genai.Content) string {
	var output string
	for _, part := range content.Parts {
		if text, ok := part.(genai.Text); ok {
			output += string(text)
		}
	}
	return output
}
