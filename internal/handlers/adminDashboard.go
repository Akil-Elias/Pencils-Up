package handlers

import (
	"html/template"
	"net/http"
)

func AdminDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("../internal/views/admin_dashboard.html"))
	templ.Execute(w, nil)
}
