package service

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func TranslateMenu(ctx context.Context, id string, imageBytes []byte, imageFormat string) (*dto.MenuTranslationResponse, *fail.Fail) {
	user, err := repository.GetUserById(ctx, id)

	if err != nil {
		return nil, &fail.UserNotFound
	}

	boundingBoxResult, err := gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash").
		WithImage(imageBytes, imageFormat).
		WithTemperature(0.2).
		WithPrompt(createBoundingBoxPrompt()).
		ExpectResponseType("application/json").
		Generate()

	if err != nil {
		return nil, &fail.FailedTranslationGeneration
	}

	output := &domain.Menu{}

	_, err = gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash-001").
		WithImage(imageBytes, imageFormat).
		WithTemperature(0.2).
		WithPrompt(createMenuTranslationPrompt(boundingBoxResult, user.Language)).
		ExpectStructuredOutput(output).
		Generate()

	if err != nil {
		return nil, &fail.FailedTranslationGeneration
	}
	return dto.FromMenu(output), nil
}

func createBoundingBoxPrompt() string {
	return `Detect 2D bounding boxes around all orderable elements in the menu image—such as menu item names, options, sub-items, etc.—but exclude any prices.
For each detected box, output:
	- "label: the exact text content inside the box  
	- "box_2d": [ymin, xmin, ymax, xmax]`
}

func createMenuTranslationPrompt(boundingBoxData string, language string) string {
	return fmt.Sprintf(`You are a menu-image translator AI.

Your input will consist of:

1. A menu image.
2. A JSON array of objects, each with:
	- box_2d: [y1, x1, y2, x2]
	- label: the raw text detected in that box

Your task: Analyze the image and the bounding boxes to determine which labels correspond to main menu items (dishes or drinks) and which correspond to options.

For each main menu item, output a JSON object with:
- label: the original label text. the exact "lable" value from input JSON Data. DO NOT MODIFY.
- box_2d: the original coordinates. the exact "box_2d" value from input JSON Data. DO NOT MODIFY.
- price: price of the menu item.
- originalItemName: a fully descriptive name of the dish or drink in its original language, inferred from visual and contextual clues. It must:
	- Clearly state the exact food name
	- If it’s a combo menu, specify which dish the combo belongs to
	- Reflect the grouping and layout seen in the image
- translatedItemName: the %s translation of that inferred name
- availableOptions: an array of translated option names that apply to this item. Use an empty array if no options apply.

Use image context to determine which options belong to each item, and include any add-ons or sides detected.
For each option label, assign it under the most logical parent item’s availableOptions, translating it into natural %s.
Do not include any unorderable menu items.
Option items must never be output as parent items.
If a detected label is ambiguous, use image context to infer a more descriptive menu name.
Return only the JSON array of main menu item objects as described—no additional text or commentary.

Following is JSON Array Object
%s`, language, language, boundingBoxData)
}
