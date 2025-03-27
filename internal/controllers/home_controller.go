package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeHandler handles requests to the homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Build the path to the template file.
	tmplPath := filepath.Join("internal", "views", "templates", "index.html")

	// Parse the template file.
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Data to pass to the template.
	data := map[string]interface{}{
		"Title": "Welcome to the ARGO CMS",
	}

	// Execute the template with data.
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
