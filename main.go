package main

import (
	"log"
	"net/http"

	"asciiArtWeb/functions"
)

func main() {
	// Route handlers
	http.HandleFunc("/", functions.HandlerIndex)
	http.HandleFunc("/ascii-art", functions.HandlerPost)
	http.HandleFunc("/static/", functions.HandleStatic)

	// Start server
	log.Println("Server running on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
