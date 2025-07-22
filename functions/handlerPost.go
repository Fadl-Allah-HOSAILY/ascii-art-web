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

	// Récupérer les données du formulaire
	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text, checkText := r.PostForm["text"]
	banner, checkBanner := r.PostForm["banner"]
	if !checkText{
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if !checkBanner{
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if len(text[0]) > 200 {
		ErrorHandler(w, "input too long", http.StatusBadRequest)
		return
	}
	// Ajouter le dossier et extension pour trouver le fichier correctement
	bannerPath := "banners/" + banner[0] + ".txt"

	asciiMap := ReadAsciiBanner(bannerPath)
	if asciiMap == nil {
		ErrorHandler(w, "Error in banner reading", http.StatusInternalServerError)
		return
	}

	asciiResult := AsciiRepresentation(text[0], asciiMap)

	data := map[string]string{
		"Ascii": asciiResult,
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, "Error", http.StatusInternalServerError)
	}
}
