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

	temperature := float32(0.5)
	topp := float32(0.8)
	topk := int32(1)

	boundingBoxResult, err := gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.5-pro-exp-03-25").
		WithImage(imageBytes, imageFormat).
		WithTemperature(&temperature).
		WithTopK(&topk).
		WithTopP(&topp).
		WithPrompt(createBoundingBoxPrompt()).
		// ExpectResponseType("application/json").
		Generate()

	if err != nil {
		return nil, &fail.FailedTranslationGeneration
	}

	output := &domain.Menu{}

	_, err = gemini.GeminiRequestBuilder().
		WithContext(ctx).
		WithModel("gemini-2.0-flash-001").
		WithImage(imageBytes, imageFormat).
		WithTemperature(&temperature).
		WithTopK(&topk).
		WithPrompt(createMenuTranslationPrompt(boundingBoxResult, user.Language)).
		ExpectStructuredOutput(output).
		Generate()

	if err != nil {
		return nil, &fail.FailedTranslationGeneration
	}
	return dto.FromMenu(output), nil
}

func createBoundingBoxPrompt() string {
	return `Extract the text from the image

Return just box_2d which will be location of detected text areas + label`
	// `Role
	// You are a meticulous menu-image annotator. Your job is to mark every orderable element—menu item names, options, sub-items—while ignoring price text. Output must be precise, concise, and limited to the required JSON.

	// Task
	// 1. Input
	//    • one menu image

	// 2. Detect a 2-D bounding box for each piece of text that represents an orderable element:
	//    • main item names
	//    • options, add-ons, sizes, flavors, sub-items
	//    (Skip any text that shows only a price.)

	// 3. For every detected box, create an object with the keys below and
	//    return all objects as a JSON array.

	//    label    the exact text content inside the box (do not alter the text)
	//    box_2d   [ymin, xmin, ymax, xmax] coordinates of the box

	// Guidelines
	// • Before producing the final answer, think step by step and verify each candidate box:
	//   - Confirm the text is orderable.
	//   - Confirm the text appears within the specified [ymin, xmin, ymax, xmax] box in the image.
	//   - Confirm the text is not solely a price or currency symbol.
	//   - Confirm the coordinate order is [ymin, xmin, ymax, xmax].
	// • Include a box whenever the text can be selected during ordering.
	// • Omit boxes whose text is solely a price or currency symbol.
	// • Use the coordinate order [ymin, xmin, ymax, xmax] consistently.
	// • Return only the JSON array—no additional text.

	// Output format
	// [
	//
	//	{ "label": "<text>", "box_2d": [ymin, xmin, ymax, xmax] },
	//	…
	//
	// ]`
}

func createMenuTranslationPrompt(boundingBoxData string, language string) string {
	return fmt.Sprintf(`- Role
You are a meticulous bilingual culinary data-extractor and menu translator. Your responses are concise, professional, and focused on accurately mapping menu items and their options into structured data.


- Task
1. Input
   • one menu image  
   • a JSON array of objects:
     { "box_2d": [y1, x1, y2, x2], "label": "<raw text>" }

2. Examine the image and decide for each label whether it represents
   • a main menu item that can be ordered, or
   • an option such as an add-on, side, size, or flavor.

3. For every main menu item, create an object with the keys below and
   return all objects together as a JSON array.

   label               use the “label” value exactly as it appears in the input
   box_2d              use the “box_2d” value exactly as it appears in the input
   price               price shown for the item
   originalItemName    a fully descriptive name in the original language
   translatedItemName  a natural %s translation of originalItemName  
                       (provide a meaningful equivalent rather than a phonetic transliteration)
   availableOptions    an array of %s option names for this item
                       (use [] if the item has no options)

- Guidelines
• Place each option label inside the availableOptions array of its parent item.  
• Include only orderable menu items in the output.  
• When a label could belong to several items, think step by step and attach it to the most plausible parent based on layout and visual grouping.  
• Express combo relationships and other context clearly in originalItemName.  
• Translate every name into natural %s, choosing meaningful wording over phonetic transliteration.  

- Few-shot examples
-----------------

Example 1
Input boxes
[
  { "box_2d": [0,5,20,100], "label": "스타벅스" },
  { "box_2d": [5,5,20,100], "label": "에스프레소 류" },
  { "box_2d": [10,10,40,200], "label": "스타벅스 아메리카노" },
  { "box_2d": [10,210,40,300], "label": "₩4,500" },
  { "box_2d": [50,10,80,200], "label": "Hot / Iced" }
]

Expected output
[
  {
    "label": "스타벅스 아메리카노",
    "box_2d": [10,10,40,200],
    "price": 4500,
    "originalItemName": "아메리카노",
    "translatedItemName": "Americano",
    "availableOptions": ["Hot", "Iced"]
  }
]

Example 2
Input boxes
[
  { "box_2d": [5,5,30,200], "label": "BBQ 치킨 연세대점" },
  { "box_2d": [10,10,40,250], "label": "황금올리브 후라이드 세트" },
  { "box_2d": [10,260,40,330], "label": "₩22,000" },
  { "box_2d": [50,10,80,120], "label": "기본맛" },
  { "box_2d": [50,130,80,250], "label": "매콤한맛" },
  { "box_2d": [90,10,120,200], "label": "소스 추가" }
]

Expected output
[
  {
    "label": "황금올리브 후라이드 세트",
    "box_2d": [10,10,40,250],
    "price": 22000,
    "originalItemName": "후라이드 치킨 세트",
    "translatedItemName": "Fried Chicken Combo",
    "availableOptions": ["Original", "Spicy", "Extra Sauce"]
  }
]

Example 3
Input boxes
[
  { "box_2d": [5,120,10,200], "label": "John Pizzeria" },
  { "box_2d": [10,10,40,150], "label": "존의 마르게리따 피자" },
  { "box_2d": [10,160,40,220], "label": "₩16,000" },
  { "box_2d": [50,10,80,150], "label": "치즈 추가" },
  { "box_2d": [10,230,40,380], "label": "페페로니 피자" },
  { "box_2d": [10,390,40,440], "label": "₩17,000" },
  { "box_2d": [50,230,80,380], "label": "존의 페페로니 추가" }
]

Expected output
[
  {
    "label": "존의 마르게리따 피자",
    "box_2d": [10,10,40,150],
    "price": 16000,
    "originalItemName": "마르게리따 피자",
    "translatedItemName": "Margherita Pizza",
    "availableOptions": ["Add Cheese"]
  },
  {
    "label": "존의 페페로니 피자",
    "box_2d": [10,230,40,380],
    "price": 17000,
    "originalItemName": "페페로니 피자",
    "translatedItemName": "Pepperoni Pizza",
    "availableOptions": ["Extra Pepperoni"]
  }
]

- Output format
Respond with the JSON array of main-menu objects as your entire reply.

Following is the bounding-box JSON list:
%s`, language, language, language, boundingBoxData)
}
