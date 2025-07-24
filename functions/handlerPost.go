package functions

import (
	"net/http"
	"text/template"
)

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Error, method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text:= r.FormValue("text")
	Banner:= r.FormValue("banner")
	
	if len(text) > 1000 {
		ErrorHandler(w, "input too long", http.StatusBadRequest)
		return
	}
	// Ajouter le dossier et extension pour trouver le fichier correctement
	bannerPath := "banners/" + Banner + ".txt"

	asciiMap := ReadAsciiBanner(bannerPath)
	if asciiMap == nil {
		ErrorHandler(w, "Error in banner reading", http.StatusInternalServerError)
		return
	}

	asciiResult := AsciiRepresentation(text, asciiMap)

	data := map[string]string{
		"Ascii": asciiResult,
	}
	tmpl,err:= template.ParseFiles("templates/index.html")
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, "Error", http.StatusInternalServerError)
	}
}
