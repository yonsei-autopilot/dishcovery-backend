package controller

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var menuTranslationPageTmpl = template.Must(template.ParseFiles(
	filepath.Join("resources", "templates", "menu_translation_page.html"),
))

func renderMenuTranslationPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := menuTranslationPageTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "템플릿 렌더링 오류", http.StatusInternalServerError)
	}
}
