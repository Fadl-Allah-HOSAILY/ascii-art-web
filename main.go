package main

import (
	"log"
	"net/http"

	"asciiArtWeb/functions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functions.HandlerIndex)
	mux.HandleFunc("/ascii-art", functions.HandlerPost)
	log.Println("Server running on: http://localhost:8080")
	mux.HandleFunc("/static/", functions.HandleStatic)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Erreur serveur :", err)
	}
}
