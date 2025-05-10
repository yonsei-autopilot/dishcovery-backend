package dto

import answer "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/gemini"

type MenuOrderResponse struct {
	OrderInUserLanguage                     string `json:"orderInUserLanguage"`
	OrderInForeignLanguage                  string `json:"orderInForeignLanguage"`
	InquiryForDislikeFoodsInUserLanguage    string `json:"inquiryForDislikeFoodsInUserLanguage"`
	InquiryForDislikeFoodsInForeignLanguage string `json:"inquiryForDislikeFoodsInForeignLanguage"`
	OrderAudioBase64                        string `json:"orderAudioBase64"`
	InquiryAudioBase64                      string `json:"inquiryAudioBase64"`
}

func FromMenuOrderInfos(menuOrderAnswer *answer.MenuOrderAnswer, orderAudioBase64 string, inquiryAudioBase64 string) *MenuOrderResponse {
	return &MenuOrderResponse{
		OrderInUserLanguage:                     menuOrderAnswer.OrderInUserLanguage,
		OrderInForeignLanguage:                  menuOrderAnswer.OrderInForeignLanguage,
		InquiryForDislikeFoodsInUserLanguage:    menuOrderAnswer.InquiryForDislikeFoodsInUserLanguage,
		InquiryForDislikeFoodsInForeignLanguage: menuOrderAnswer.InquiryForDislikeFoodsInForeignLanguage,
		OrderAudioBase64:                        orderAudioBase64,
		InquiryAudioBase64:                      inquiryAudioBase64,
	}
}
