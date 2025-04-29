package service

import (
	"context"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
)

type Menu struct {
	Items []Item `json:"items" genai:"description=Menu items list;required"`
}

type Item struct {
	Name        string  `json:"name" genai:"description=Item name;required"`
	Description string  `json:"description" genai:"description=Description of the item. Do not include any item name or price info.;required"`
	Price       float32 `json:"price" genai:"description=Price of item"`
}

func ExplainMenu(imageBytes []byte, imageFormat string) (string, error) {
	ctx := context.Background()

	output := &Menu{}

	return gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithImage(imageBytes, imageFormat).
		WithPrompt("Given a menu, describe each item in Korean. Include the dish name, price, and a detailed explanation of its ingredients, flavors, and characteristics.").
		ExpectStructuredOutput(output).
		Generate()
}
