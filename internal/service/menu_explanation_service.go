package service

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
)

func ExplainMenu(imageBytes []byte, imageFormat string) (string, error) {
	return gemini.Send(
		imageBytes,
		imageFormat,
		"Explain this restaurant menu image in English.",
	)
}
