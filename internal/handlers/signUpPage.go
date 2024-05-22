package handlers

import (
	"html/template"
	"net/http"
)

func SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("../internal/views/sign_up.html"))
	templ.Execute(w, nil)
}
