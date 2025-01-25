package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the static JSON file from the data directory
	http.Handle("/", http.FileServer(http.Dir("data")))

	// Start the local JSON server
	log.Println("JSON server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
