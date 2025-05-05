package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/domain"

type MenuExplanationResponse struct {
	Items []MenuItemResponse `json:"items"`
}

type MenuItemResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func FromMenu(menu *domain.Menu) *MenuExplanationResponse {
	items := make([]MenuItemResponse, len(menu.Items))
	for i, item := range menu.Items {
		items[i] = MenuItemResponse{
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		}
	}
	return &MenuExplanationResponse{Items: items}
}
