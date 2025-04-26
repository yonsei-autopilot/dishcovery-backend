package service

import (
	"strings"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
)

func ExplainMenu(imageBytes []byte, imageFormat string) (string, error) {
	return gemini.Send(
		imageBytes,
		imageFormat,
		"Explain this restaurant menu image in English.",
	)
}

func extractJSONFromMarkdownBlock(input string) string {
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "```json")
	input = strings.TrimPrefix(input, "```")
	input = strings.TrimSuffix(input, "```")
	return strings.TrimSpace(input)
}

func BoundMenu(imageBytes []byte, imageFormat string) (string, error) {
	output, err := gemini.Send(
		imageBytes,
		imageFormat,
		"Detect each menu items. Outputlist where each entry contains the 2D bounding box in \"box_2d\" and menu name in \"label\". Output must be JSON format. Do not include any prices information, Menu name must be translated into Korean.",
	)
	return extractJSONFromMarkdownBlock(output), err
}
