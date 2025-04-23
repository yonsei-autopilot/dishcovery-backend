package controller

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var menuExplainPageTmpl = template.Must(template.ParseFiles(
	filepath.Join("resources", "templates", "menu_explanation_page.html"),
))

func renderMenuExplanationPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := menuExplainPageTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "템플릿 렌더링 오류", http.StatusInternalServerError)
	}
}
