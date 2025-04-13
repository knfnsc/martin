package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
		if tmpl == nil {
			tmpl = template.Must(tmpl.ParseGlob("views/*.html"))
		}
}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Martin",
	}

	tmpl.ExecuteTemplate(w, "index.html", data)
}

func ShowEditor(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Martin — Editor",
	}

	tmpl.ExecuteTemplate(w, "editor.html", data)
}
