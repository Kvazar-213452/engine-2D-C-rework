package main_

import (
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title string
	Body  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "Моя перша веб-сторінка",
		Body:  "Вітаю на моїй веб-сторінці!",
	}

	// Завантаження шаблонів
	tmpl, err := template.ParseFiles("front/page/header.pug", "front/page/index.pug")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Виконання шаблону
	err = tmpl.ExecuteTemplate(w, "index.pug", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
