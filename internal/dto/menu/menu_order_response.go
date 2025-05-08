package dto

import answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"

type MenuOrderResponse struct {
	OrderInUserLanguage                     string `json:"orderInUserLanguage"`
	OrderInForeignLanguage                  string `json:"orderInForeignLanguage"`
	InquiryForDislikeFoodsInUserLanguage    string `json:"inquiryForDislikeFoodsInUserLanguage"`
	InquiryForDislikeFoodsInForeignLanguage string `json:"inquiryForDislikeFoodsInForeignLanguage"`
}

func FromMenuOrderAnswer(menuOrderAnswer *answer.MenuOrderAnswer) *MenuOrderResponse {
	return &MenuOrderResponse{
		OrderInUserLanguage:                     menuOrderAnswer.OrderInUserLanguage,
		OrderInForeignLanguage:                  menuOrderAnswer.OrderInForeignLanguage,
		InquiryForDislikeFoodsInUserLanguage:    menuOrderAnswer.InquiryForDislikeFoodsInUserLanguage,
		InquiryForDislikeFoodsInForeignLanguage: menuOrderAnswer.InquiryForDislikeFoodsInForeignLanguage,
	}
}
