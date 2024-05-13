package handlers

import (
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("../internal/views/login.html"))
	templ.Execute(w, nil)
}
