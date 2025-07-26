package functions

import (
	"html/template"
	"net/http"
)

// Index page handler
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Error, page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, "Error, method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Error, internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
