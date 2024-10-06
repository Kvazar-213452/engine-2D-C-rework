package main_

import (
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"front/add/engene/modal.pug",
		"front/add/engene/inspector.pug",
		"front/index.pug",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageVariables := map[string]interface{}{
		"Title": "Моя перша веб-сторінка",
		"Body":  "Вітаю на моїй веб-сторінці!",
	}

	err = tmpl.ExecuteTemplate(w, "index.pug", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
