package answer

type MenuOrderAnswer struct {
	OrderInUserLanguage                     string `json:"orderInUserLanguage" genai:"description=The order statement which is in user language.;required"`
	OrderInForeignLanguage                  string `json:"OrderInForeignLanguage" genai:"description=The order statement which is in foreign language;required"`
	InquiryForDislikeFoodsInUserLanguage    string `json:"inquiryForDislikeFoodsInUserLanguage" genai:"description=The inquiry for dislike foods which is in user language.;required"`
	InquiryForDislikeFoodsInForeignLanguage string `json:"inquiryForDislikeFoodsInForeignLanguage" genai:"description=The inquiry for dislike foods which is in foreign language.;required"`
}
