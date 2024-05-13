package handlers

import (
	"net/http"
	"strings"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// templ := template.Must(template.ParseFiles("../internal/views/home.html"))
	// templ.Execute(w, nil)
	path := "../internal/views/home.html"
	// Check if the request is for an image file
	if strings.HasSuffix(r.URL.Path, ".jpg") || strings.HasSuffix(r.URL.Path, ".jpeg") || strings.HasSuffix(r.URL.Path, ".png") {
		// Serve image file
		http.ServeFile(w, r, "../internal"+r.URL.Path)
		return
	}

	// For non-image files (e.g., HTML), use http.ServeFile
	http.ServeFile(w, r, path)
}
