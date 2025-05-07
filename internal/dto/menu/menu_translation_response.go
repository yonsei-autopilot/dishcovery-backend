package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/domain"

type MenuTranslationResponse struct {
	Items []MenuItemResponse `json:"items"`
}

type MenuItemResponse struct {
	OriginalItemName   string   `json:"originalItemName"`
	TranslatedItemName string   `json:"translatedItemName"`
	Label              string   `json:"label"`
	AvailableOptions   []string `json:"availableOptions"`
	Price              float32  `json:"price"`
	BoundingBox        []int    `json:"boundingBox"`
}

func FromMenu(menu *domain.Menu) *MenuTranslationResponse {
	items := make([]MenuItemResponse, len(menu.Items))
	for i, item := range menu.Items {
		items[i] = MenuItemResponse{
			TranslatedItemName: item.TranslatedItemName,
			OriginalItemName:   item.OriginalItemName,
			Label:              item.Label,
			AvailableOptions:   item.AvailableOptions,
			BoundingBox:        item.BoundingBox,
			Price:              item.Price,
		}
	}
	return &MenuTranslationResponse{Items: items}
}
