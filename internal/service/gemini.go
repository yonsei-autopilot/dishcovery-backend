package service

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/config"
	"google.golang.org/api/option"
)

func SendToGemini(imageBytes []byte) (string, error) {
	ctx := context.Background()

	apiKey, err := config.GetEnv("GEMINI_API_KEY")
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	prompt := genai.Text("Please explain this restaurant menu image in English:")

	res, err := model.GenerateContent(ctx, prompt, genai.ImageData("image/jpeg", imageBytes))
	if err != nil {
		return "", fmt.Errorf("Gemini generation failed: %w", err)
	}

	// Validate response structure
	if len(res.Candidates) == 0 || len(res.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("Gemini returned an unexpected empty response")
	}

	textPart, ok := res.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		log.Printf("Gemini response part is not text: %#v", res.Candidates[0].Content.Parts[0])
		return "", fmt.Errorf("Gemini response part is not valid text")
	}

	return string(textPart), nil
}
