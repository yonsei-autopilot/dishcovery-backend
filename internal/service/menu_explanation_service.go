package service

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

func ExplainMenu(imageBytes []byte, imageFormat string) (string, error) {
	return util.AskGemini(
		imageBytes,
		imageFormat,
		"Explain this restaurant menu image in English.",
	)
}
