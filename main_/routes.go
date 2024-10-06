package main_

import (
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/page/index.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, map[string]string{"Title": "title", "Body": "body"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
