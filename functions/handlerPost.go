package functions

import (
	"net/http"
	"text/template"
)

// Post method handler
func HandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Error, method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > 1000 {
		ErrorHandler(w, "input too long", http.StatusBadRequest)
		return
	}

	bannerPath := "banners/" + banner + ".txt"
	asciiMap := ReadAsciiBanner(w,*r,bannerPath)

	asciiResult := AsciiRepresentation(text, asciiMap)

	data := map[string]string{"Ascii": asciiResult}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, "Error", http.StatusInternalServerError)
	}
}
