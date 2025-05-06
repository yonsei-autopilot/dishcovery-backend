package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/domain"

type MenuExplanationResponse struct {
	ImageLinks      []string `json:"imageLinks"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Ingredients     string   `json:"ingredients"`
	WhatToBeCareful string   `json:"whatToBeCareful"`
}

func NewMenuExplanationResponse(imageSearchResult *ImageSearchResult, menuExplanation *domain.MenuExplanation) *MenuExplanationResponse {
	imageLinks := make([]string, len(imageSearchResult.Items))
	for i, item := range imageSearchResult.Items {
		imageLinks[i] = item.Link
	}
	return &MenuExplanationResponse{
		ImageLinks:      imageLinks,
		Name:            menuExplanation.Name,
		Description:     menuExplanation.Description,
		Ingredients:     menuExplanation.Ingredients,
		WhatToBeCareful: menuExplanation.WhatToBeCareful,
	}
}
