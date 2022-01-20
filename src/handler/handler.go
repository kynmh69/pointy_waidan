package handler

import (
	"html/template"
	"log"
	"net/http"
	"pointy/control"
	"pointy/model"
	"time"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	articls := control.GetAllData()

	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	if err := t.Execute(w, struct {
		Title   string
		Message string
		Time    time.Time
		As      *[]*model.Article
	}{
		Title:   "テストページ",
		Message: "こんにちは！",
		Time:    time.Now(),
		As:      articls,
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func LoadTemplate(name string) *template.Template {
	t, err := template.ParseFiles("template/" + name + ".html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}
